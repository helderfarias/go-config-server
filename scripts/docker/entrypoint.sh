#!/usr/bin/env sh

set -eC

# usage: file_env VAR [DEFAULT]
#    ie: file_env 'XYZ_NAME' 'example'
# (will allow for "$XYZ_NAME_FILE" to fill in the value of
#  "$XYZ_NAME" from a file, especially for Docker's secrets feature)
file_env() {
	local var="$1"
	local fileVar="${var}_FILE"
	local def="${2:-}"

  eval var_e=\$$var
  eval file_var_e=\$$fileVar

  if [ ! -z "${var_e:-}" ] && [ ! -z "${file_var_e:-}" ]
  then
	  echo >&2 "error: both $var and $fileVar are set (but are exclusive)"
		exit 1
	fi

	local val="$def"

	if [ -n "${var_e:-}" ]
  then
		val="${var_e}"
	elif [ -n "${file_var_e:-}" ]
  then
		val="$(cat $file_var_e)"
	fi

	export "$var"="$val"
  unset "$fileVar"
}

file_env 'CONFIG'

/bin/gcs -config ${CONFIG}
