# max_mtime

Given a list of files as arguments returns the maximum mtime of the files to stdout

By rights this shouldn't exist. `stat` can return this information and a simple `bash` script would be able to extract the maximum value

Until you find that `stat` is not quite as standard across the various Unix variants. As a Mac and Linux user I keep stumbling into how BSD is just different enough from Linux for this to be a persistent issue

So here we are
