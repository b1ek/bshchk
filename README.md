# bshchk
a runtime bash dependency checker. this is useful when you want to assure that ALL external commands you have used will be present at runtime.

## usage
```sh
bshchk source.sh with_bshchk.sh
bshchk '' with_bshchk.sh # will read from stdin
bshchk # will read from stdin, and write to stdin
```

tags
```sh
# To explicitly add curl:
#bshchk:add-cmd curl

# To disable checking for one curl:
#bshchk:ignore-cmd curl

# You can add/ignore multiple commands at once:
#bshchk:add-cmd curl wget
```


## license
```
bshchk - a bash runtime dependency checker
Copyright (C) 2024  blek! <me@blek.codes>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, version 3  of the License,  and not  a
later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```