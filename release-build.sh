#!/usr/bin/env bash
unset SCRIPT_ROOT
readonly SCRIPT_ROOT=$(
  unset CDPATH
  script_root=$(dirname "${BASH_SOURCE}")
  cd "${script_root}"
  pwd
)
if [ -d "${SCRIPT_ROOT}/submodules/flogo-cicd" ]; then
  rm -rf ${SCRIPT_ROOT}/submodules/flogo-cicd
  git submodule update --init --remote -- submodules/flogo-cicd
  source ${SCRIPT_ROOT}/submodules/flogo-cicd/scripts/init.sh
  # Build flogo/flogo-contrib docker image
  pushd ${SCRIPT_ROOT}
  docker::build_and_push flogo/flogo-contrib
  # docker::copy_tag_and_push flogo/flogo-contrib pdhar/flogo-contrib
  popd
fi
