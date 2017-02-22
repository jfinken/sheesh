# sheesh

sheesh returns the first 8 digits of the hexidecimal digest of the md5sum of the input string.  The input string can be piped to stdin or received as argument one.


    $sheesh fooie
    0b2d454d

    $echo "fooie" | sheesh
    0b2d454d
