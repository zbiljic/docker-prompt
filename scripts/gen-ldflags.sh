#!/usr/bin/env bash

_init() {
    ##
    ## IMPORTANT: DO NOT change order
    ##
    _build_time=$(buildTime)
    _version=$(version)
    _commit_id=$(commitID)
    _short_commit_id=${_commit_id:0:7}
}

main() {
    ld_flags_str="-s -w"
    ld_flags_str="$ld_flags_str -X github.com/zbiljic/docker-prompt/cmd.Version=$_version"
    ld_flags_str="$ld_flags_str -X github.com/zbiljic/docker-prompt/cmd.BuildTime=$_build_time"
    ld_flags_str="$ld_flags_str -X github.com/zbiljic/docker-prompt/cmd.CommitID=$_commit_id"
    ld_flags_str="$ld_flags_str -X github.com/zbiljic/docker-prompt/cmd.ShortCommitID=$_short_commit_id"
    echo "$ld_flags_str"
}

# buildTime returns current time in the format '%Y-%m-%dT%H-%M-%SZ'
buildTime() {
    time_str=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
    echo "${time_str//:/-}"
}

version() {
    version="dev"

    # extract Git tag for currenct commit
    git_version=$(git describe --tags --exact-match --abbrev=0 2>/dev/null)
    git_version="${git_version// }"

    if [[ ! -z "$git_version" ]]; then
        # release version
        version="$git_version"
    else
        # development version
        version="$version.$_build_time"
    fi

    echo "$version"
}

# commitID returns the abbreviated commit-id hash of the last commit
commitID() {
    # check if provided through environment variable (by build system)
    last_commit=$(printenv LASTCOMMIT)
    last_commit="${last_commit// }"
    if [[ ! -z "$last_commit" ]]; then
        echo "$last_commit"
        return
    fi

    last_commit="$(git rev-parse HEAD)"
    last_commit="${last_commit// }"
    echo "$last_commit"
}

_init && main || exit 1
