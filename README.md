# go-tree-view
Simple representation of tree view for file and folders written GoLang

This is very basic tree view, simple to the Unix (tree).

Here is a simple preview:

C:\Users\stefkoff\
|&nbsp;&nbsp;&nbsp;&nbsp;Favorites\
|&nbsp;&nbsp;&nbsp;&nbsp;|---Bing.url\
|&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;Links\
|&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;|---desktop.ini\
|&nbsp;&nbsp;&nbsp;&nbsp;|---desktop.ini\
|&nbsp;&nbsp;&nbsp;&nbsp;Saved Games\
|&nbsp;&nbsp;&nbsp;&nbsp;|---desktop.ini\
|&nbsp;&nbsp;&nbsp;&nbsp;www\
|&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;yaml-cpp\

Currently the programs supports te folowing options:
 - -d Max Deep - Maximum deep to go inside a folder. Default is 0 which means to go to the end
 - -h Show hidden file (NOT IMPLEMENTED YET)

# TODO
1. Add different colors for the different file modes
2. implement the hidden files options (Windows?)
3. ...