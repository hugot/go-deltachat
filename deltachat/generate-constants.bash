#!/bin/bash
##
# Generate golang mapping to deltachat constants

here="$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")"
destination="$here/constants.go"

{
    printf 'package deltachat\n\n'

    printf '// #include <deltachat.h>
import "C"\n'

    printf 'const (\n'
    
    while read -r define const_name rest; do
        if [[ "$define" == '#define' ]]; then
            if case "$const_name" in
                   DC_EVENT_DATA* | DC_EVENT_RETURNS*) false;;
                   DC_EVENT_*) :;;
                   DC_LP_*) :;;
                   DC_MSG_*) :;;
                   DC_GCL_*) :;;
                   DC_QR_*) :;;
                   DC_CHAT_*) :;;
                   DC_STATE_*) :;;
                   DC_CONTACT_*) :;;
                   DC_EMPTY_*) :;;
                   DC_STR_*) :;;
                   DC_CERTCK_*) :;;
                   DC_EMPTY_*) :;;
                   DC_PROVIDER_STATUS_*) :;;
                   *) false;;
               esac
            then
                printf '%s = int(C.%s)\n' "$const_name" "$const_name"
            fi
        fi
    done < "$here"/../deltachat-ffi/include/deltachat.h

    printf ')\n'
} > "$destination"

gofmt -w "$destination"
