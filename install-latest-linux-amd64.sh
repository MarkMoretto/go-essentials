#!/bin/bash
set -e

# Install latest Go version on Debian/Ubuntu-based distros
# Note: Requires curl or wget.

# TODO: Remove downloaded file.

_GO_LATEST_VER=
_GO_BIN_FILENANE=
_GO_BIN_DL_FILEPATH=
set-latest-go() {
    _GO_LATEST_VER=$(curl "https://go.dev/dl/?mode=json" | jq '.[0].version')
}

dl-latest-go() {
    # Set global variables.
    set-latest-go
    # Trim quotation marks from each end of variable value.
    _GO_BIN_FILENAME="${_GO_LATEST_VER:1:${#_GO_LATEST_VER}-2}.linux-amd64.tar.gz"
    # Assumes default download destination of $HOME/Downloads
    _GO_BIN_DL_FILEPATH="${HOME}/Downloads/${_GO_BIN_FILENAME}"
    # Check if already downloaded.  If not, download it.
    if [[ $(ls -l "${_GO_BIN_DL_FILEPATH}" | wc -l) = 0 ]]; then
        local _GO_DL_URL= "https://go.dev/dl/${_GO_BIN_FILENAME}"
        printf 'Retrieving %s\n' "${_GO_BIN_FILENANE}"
        wget -O "${_GO_BIN_DL_FILEPATH}" "${_GO_DL_URL}"
        # Comment-out above and uncomment below if you would like to use curl instead.
        #curl -O "${_GO_DL_URL}"
    else
        printf 'Latest Go version already downloaded to: %s\n' "${_GO_BIN_DL_FILEPATH}"
    fi
}

install-latest-go() {
    # Install latest Go version from ~/Downloads folder.
    # See: https://go.dev/doc/install
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf "${_GO_BIN_FILENAME}"
}

# Run functions in succession.
dl-latest-go && install-latest-go
