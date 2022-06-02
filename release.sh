#! /bin/bash

# Make it compatible with node.js
npm init -y > /dev/null

# Read the file version
version=`cat version.txt`

# Get the pacakge version.
existingVersion=`grep -w "version" package.json | cut -c 15-19`

# Only update if there is a change
if [ "$version" != "$existingVersion" ]; then
    # update the version of the package.
    # Avoid creating the tag for the version
    npm --no-git-tag-version version $version 
fi

# Add the files that needs to be in the pacakge.
sed -i -e '23s_}_,"files":["contracts/*", "transactions/*"]}_' package.json

# Remove the backup file.
FILE=./package.json-e
if test -f "$FILE"; then 
    rm package.json-e
fi