"""
Update everything in this project to the latest aas-core-meta.

Git is expected to be installed.
"""

from __future__ import annotations

import argparse
import os
import pathlib
import re
import shutil
import subprocess
import sys
import tempfile
import time
from typing import Optional, List, Callable, AnyStr, Sequence

# noinspection RegExpSimplifiable
AAS_CORE_META_DEPENDENCY_RE = re.compile(
    r"aas-core-meta@git\+https://github.com/aas-core-works/aas-core-meta@([a-fA-F0-9]+)"
)

# noinspection RegExpSimplifiable
AAS_CORE_CODEGEN_DEPENDENCY_RE = re.compile(
    r"aas-core-codegen@git\+https://github.com/aas-core-works/aas-core-codegen@([a-fA-F0-9]+)"
)


def _make_sure_no_changed_files(
    repo_dir: pathlib.Path, expected_branch: str
) -> Optional[int]:
    """
    Make sure that no files are modified in the given repository.

    Return exit code if something is unexpected.
    """
    diff_name_status = subprocess.check_output(
        ["git", "diff", "--name-status", expected_branch],
        cwd=str(repo_dir),
        encoding="utf-8",
    ).strip()

    if len(diff_name_status.splitlines()) > 0:
        print(
            f"The following files are modified "
            f"compared to branch {expected_branch!r} in {repo_dir}:\n"
            f"{diff_name_status}\n"
            f"\n"
            f"Please stash the changes first before you update to aas-core-meta.",
            file=sys.stderr,
        )
        return 1

    return None


def _update_pyproject_toml(
    our_repo: pathlib.Path, aas_core_meta_revision: str, aas_core_codegen_revision: str
) -> None:
    """Update the aas-core-meta in pyproject.toml."""
    pyproject_toml = our_repo / "_dev_scripts" / "pyproject.toml"
    text = pyproject_toml.read_text(encoding="utf-8")

    aas_core_meta_dependency = (
        f"aas-core-meta@git+https://github.com/aas-core-works/aas-core-meta"
        f"@{aas_core_meta_revision}"
    )

    text = re.sub(AAS_CORE_META_DEPENDENCY_RE, aas_core_meta_dependency, text)

    aas_core_codegen_dependency = (
        f"aas-core-codegen@git+https://github.com/aas-core-works/aas-core-codegen"
        f"@{aas_core_codegen_revision}"
    )

    text = re.sub(AAS_CORE_CODEGEN_DEPENDENCY_RE, aas_core_codegen_dependency, text)

    pyproject_toml.write_text(text, encoding="utf-8")


def _uninstall_and_install_aas_core_meta(
    our_repo: pathlib.Path, aas_core_meta_revision: str
) -> None:
    """Uninstall and install the latest aas-core-meta in the virtual environment."""
    subprocess.check_call(
        [sys.executable, "-m", "pip", "uninstall", "-y", "aas-core-meta"],
        cwd=str(our_repo),
    )

    aas_core_meta_dependency = (
        f"aas-core-meta@git+https://github.com/aas-core-works/aas-core-meta"
        f"@{aas_core_meta_revision}#egg=aas-core-meta"
    )

    subprocess.check_call(
        [sys.executable, "-m", "pip", "install", aas_core_meta_dependency],
        cwd=str(our_repo),
    )


def _uninstall_and_install_aas_core_codegen(
    our_repo: pathlib.Path, aas_core_codegen_revision: str
) -> None:
    """Uninstall and install the latest aas-core-codegen in the virtual environment."""
    subprocess.check_call(
        [sys.executable, "-m", "pip", "uninstall", "-y", "aas-core-codegen"],
        cwd=str(our_repo),
    )

    aas_core_codegen_dependency = (
        f"aas-core-codegen@git+https://github.com/aas-core-works/aas-core-codegen"
        f"@{aas_core_codegen_revision}#egg=aas-core-codegen"
    )

    subprocess.check_call(
        [sys.executable, "-m", "pip", "install", aas_core_codegen_dependency],
        cwd=str(our_repo),
    )


def _regenerate_code(
    our_repo: pathlib.Path,
    meta_repo: pathlib.Path,
) -> Optional[int]:
    """
    Call codegen script.

    Return an error code, if any.
    """
    codegen_dir = our_repo / "_dev_scripts/codegen"

    meta_model_path = meta_repo / "aas_core_meta/v3.py"

    target_dir = our_repo

    print(f"Starting to run codegen script")
    start = time.perf_counter()

    proc = subprocess.run(
        [
            sys.executable,
            "codegen.py",
            "--meta_model",
            str(meta_model_path),
            "--target",
            str(target_dir),
        ],
        cwd=str(codegen_dir),
    )

    if proc.returncode != 0:
        return proc.returncode

    duration = time.perf_counter() - start
    print(f"Generating the code took: {duration:.2f} seconds.")

    return None


def _reformat_code(our_repo: pathlib.Path) -> None:
    """Reformat the generated code."""
    print("Re-formatting the code...")

    # NOTE (empwilli):
    # We have to explicitly list all directories, otherwise gofmt fails on our
    # snippets dir.
    go_dirs = [
        "aastesting",
        "common",
        "constants",
        "getting_started",
        "jsonization",
        "reporting",
        "stringification",
        "types",
        "verification",
        "xmlization",
    ]

    # NOTE (mristin, 2023-06-24):
    # We have to call gofmt twice to replace the bullet points to its latest
    # standard.
    subprocess.check_call(["gofmt", "-w", *go_dirs], cwd=our_repo)
    subprocess.check_call(["gofmt", "-w", *go_dirs], cwd=our_repo)


def _run_tests_and_rerecord(our_repo: pathlib.Path) -> None:
    """Run the tests with the environment variables set to re-record."""
    print("Running tests & re-recording the test traces...")

    env = os.environ.copy()
    env["AAS_CORE_3_0_GOLANG_TEST_DATA_DIR"] = str(our_repo / "testdata")
    env["AAS_CORE_3_0_GOLANG_RECORD_MODE"] = "true"

    subprocess.check_call(["go", "test", "./..."], env=env, cwd=our_repo)


def _run_check(our_repo: pathlib.Path) -> None:
    """Run the linter, build and test."""
    print("Running the pre-commit checks...")

    env = os.environ.copy()
    env["AAS_CORE3_0_GOLANG_TEST_DATA_DIR"] = str(our_repo / "testdata")

    subprocess.check_call(["go", "vet", "./..."], cwd=our_repo)


def _create_branch_commit_and_push(
    our_repo: pathlib.Path,
    aas_core_meta_revision: str,
    aas_core_codegen_revision: str,
    aas_core_testgen_revision: str,
) -> None:
    """Create a feature branch, commit the changes and push it."""
    branch = (
        f"Update-to-aas-core-meta-codegen-testgen-{aas_core_meta_revision}-"
        f"{aas_core_codegen_revision}-{aas_core_testgen_revision}"
    )
    print(f"Creating the branch {branch!r}...")
    subprocess.check_call(["git", "checkout", "-b", branch], cwd=our_repo)

    print("Adding files...")
    subprocess.check_call(["git", "add", "."], cwd=our_repo)

    # pylint: disable=line-too-long
    message = f"""\
Update to aas-core-meta, codegen, testgen {aas_core_meta_revision}, {aas_core_codegen_revision}, {aas_core_testgen_revision}

We update the development requirements to and re-generate everything
with:
* [aas-core-meta {aas_core_meta_revision}],
* [aas-core-codegen {aas_core_codegen_revision}] and
* [aas-core3.0-testgen {aas_core_testgen_revision}].

[aas-core-meta {aas_core_meta_revision}]: https://github.com/aas-core-works/aas-core-meta/commit/{aas_core_meta_revision}
[aas-core-codegen {aas_core_codegen_revision}]: https://github.com/aas-core-works/aas-core-codegen/commit/{aas_core_codegen_revision}
[aas-core3.0-testgen {aas_core_testgen_revision}]: https://github.com/aas-core-works/aas-core3.0-testgen/commit/{aas_core_testgen_revision}
"""

    # pylint: enable=line-too-long

    print("Committing...")
    with tempfile.TemporaryDirectory() as tmp_dir:
        tmp_file = pathlib.Path(tmp_dir) / "commit-message.txt"
        tmp_file.write_text(message, encoding="utf-8")

        subprocess.check_call(["git", "commit", "--file", str(tmp_file)], cwd=our_repo)

    print(f"Pushing to remote {branch}...")
    subprocess.check_call(["git", "push", "-u"], cwd=our_repo)


def _get_testgen_revision(our_repo: pathlib.Path) -> str | None:
    testgen_rev_path = our_repo / "testgen_rev.txt"

    testgen_rev: str | None = None

    try:
        with testgen_rev_path.open("r") as testgen_rev_file:
            testgen_rev = testgen_rev_file.read().strip()
    except OSError as os_error:
        print(f"Cannot read testgen revision: {os_error}.")

    if testgen_rev is None:
        print("Cannot read testgen revision.")

    return testgen_rev


def main() -> int:
    """Execute the main routine."""
    this_path = pathlib.Path(os.path.realpath(__file__))
    our_repo = this_path.parent.parent

    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--aas_core_meta_repo",
        help="path to the aas-core-meta repository",
        default=str(our_repo.parent / "aas-core-meta"),
    )
    parser.add_argument(
        "--expected_aas_core_meta_branch",
        help="Git branch expected in the aas-core-meta repository",
        default="main",
    )
    parser.add_argument(
        "--aas_core_codegen_repo",
        help="path to the aas-core-codegen repository",
        default=str(our_repo.parent / "aas-core-codegen"),
    )
    parser.add_argument(
        "--expected_aas_core_codegen_branch",
        help="Git branch expected in the aas-core-meta repository",
        default="main",
    )
    parser.add_argument(
        "--expected_our_branch",
        help="Git branch expected in this repository",
        default="main",
    )

    args = parser.parse_args()

    aas_core_meta_repo = pathlib.Path(args.aas_core_meta_repo)
    expected_aas_core_meta_branch = str(args.expected_aas_core_meta_branch)

    aas_core_codegen_repo = pathlib.Path(args.aas_core_codegen_repo)
    expected_aas_core_codegen_branch = str(args.expected_aas_core_codegen_branch)

    expected_our_branch = str(args.expected_our_branch)

    # region aas-core-meta repo

    if not aas_core_meta_repo.exists():
        print(
            f"--aas_core_meta_repo does not exist: {aas_core_meta_repo}",
            file=sys.stderr,
        )
        return 1

    if not aas_core_meta_repo.is_dir():
        print(
            f"--aas_core_meta_repo is not a directory: {aas_core_meta_repo}",
            file=sys.stderr,
        )
        return 1

    aas_core_meta_branch = subprocess.check_output(
        ["git", "rev-parse", "--abbrev-ref", "HEAD"],
        cwd=str(aas_core_meta_repo),
        encoding="utf-8",
    ).strip()
    if aas_core_meta_branch != expected_aas_core_meta_branch:
        print(
            f"--expected_aas_core_meta_branch is {expected_aas_core_meta_branch}, "
            f"but got {aas_core_meta_branch} "
            f"in --aas_core_meta_repo: {aas_core_meta_repo}",
            file=sys.stderr,
        )
        return 1

    aas_core_meta_revision = subprocess.check_output(
        ["git", "rev-parse", "--short", "HEAD"],
        cwd=str(aas_core_meta_repo),
        encoding="utf-8",
    ).strip()

    # endregion

    # region aas-core-codegen repo

    if not aas_core_codegen_repo.exists():
        print(
            f"--aas_core_codegen_repo does not exist: {aas_core_codegen_repo}",
            file=sys.stderr,
        )
        return 1

    if not aas_core_codegen_repo.is_dir():
        print(
            f"--aas_core_codegen_repo is not a directory: {aas_core_codegen_repo}",
            file=sys.stderr,
        )
        return 1

    aas_core_codegen_branch = subprocess.check_output(
        ["git", "rev-parse", "--abbrev-ref", "HEAD"],
        cwd=str(aas_core_codegen_repo),
        encoding="utf-8",
    ).strip()
    if aas_core_codegen_branch != expected_aas_core_codegen_branch:
        print(
            f"--expected_aas_core_codegen_branch is {expected_aas_core_codegen_branch}, "
            f"but got {aas_core_codegen_branch} "
            f"in --aas_core_codegen_repo: {aas_core_codegen_repo}",
            file=sys.stderr,
        )
        return 1

    aas_core_codegen_revision = subprocess.check_output(
        ["git", "rev-parse", "--short", "HEAD"],
        cwd=str(aas_core_codegen_repo),
        encoding="utf-8",
    ).strip()

    # endregion

    # region Our repo

    our_branch = subprocess.check_output(
        ["git", "rev-parse", "--abbrev-ref", "HEAD"],
        cwd=str(our_repo),
        encoding="utf-8",
    ).strip()
    if our_branch != expected_our_branch:
        print(
            f"--expected_our_branch is {expected_our_branch}, "
            f"but got {our_branch} in: {our_repo}",
            file=sys.stderr,
        )
        return 1

    # endregion

    for repo_dir, expected_branch in [
        (our_repo, expected_our_branch),
        (aas_core_meta_repo, expected_aas_core_meta_branch),
        (aas_core_codegen_repo, expected_aas_core_codegen_branch),
    ]:
        exit_code = _make_sure_no_changed_files(
            repo_dir=repo_dir, expected_branch=expected_branch
        )
        if exit_code is not None:
            return exit_code

    _update_pyproject_toml(
        our_repo=our_repo,
        aas_core_meta_revision=aas_core_meta_revision,
        aas_core_codegen_revision=aas_core_codegen_revision,
    )

    _uninstall_and_install_aas_core_meta(
        our_repo=our_repo, aas_core_meta_revision=aas_core_meta_revision
    )

    _uninstall_and_install_aas_core_codegen(
        our_repo=our_repo, aas_core_codegen_revision=aas_core_codegen_revision
    )

    exit_code = _regenerate_code(our_repo=our_repo, meta_repo=aas_core_meta_repo)
    if exit_code is not None:
        return exit_code

    _reformat_code(our_repo=our_repo)

    _run_tests_and_rerecord(our_repo=our_repo)

    _run_check(our_repo=our_repo)

    aas_core_testgen_revision = _get_testgen_revision(our_repo=our_repo)
    if aas_core_testgen_revision is None:
        return 1

    _create_branch_commit_and_push(
        our_repo=our_repo,
        aas_core_meta_revision=aas_core_meta_revision,
        aas_core_codegen_revision=aas_core_codegen_revision,
        aas_core_testgen_revision=aas_core_testgen_revision,
    )

    return 0


if __name__ == "__main__":
    sys.exit(main())
