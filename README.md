### Projet description ###
mjk - a verison control system similar to gin desigend by manas jain kuniya, it will have similar commands to git 
like git add, commit, branch, diff, etc. 
starting form scratch
based on golang


### use ###
the test_dir holdes the mjk file in it, put that file in any dir, 
cd to your directory, for which you want to creat the version control system

### current ###
developed the 
init command
commit command

### next to work on
add
diff
branch

### Documentation ###

# commands-use
1. ./mjk init
    creates a traking directory .mjk for all the version files
    it first initializes HEAD file holding the pointer to (id of) latest commit 
    and the objects folder which have all the versions of files in compressed formate
2. ./mjk commit -m "commit message"
    creates all the objects of files in the object folder

# pkg
