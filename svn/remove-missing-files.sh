svn status | grep '^!' | cut -d ' ' -f2- | xargs svn delete
