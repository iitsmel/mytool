#!/usr/bin/env bash

# Local variables
BRANCH=$(git symbolic-ref --short HEAD)

action_push() {
    # Push current base data.
    # Date=$(date)
    printf "%s\n"
    echo "Enter the commit message (use Ctrl + d to save): "
    CommitM=$(</dev/stdin)
    printf "\n"

    git init
    git add .
    git commit -m "$CommitM"
    git branch -M "$BRANCH"
    git push -u origin "$BRANCH"
}

action_pull() {
    git pull
}

action_soft() {
    # Remove Last Commit From Current Branch
    printf "%s\n"
    printf "%s\n" "Are you sure you want to remove your last commit from local and GitHub ?"
    printf "%s\n" "Enter yes or no  (use Ctrl + d to save): "
    Soft=$(</dev/stdin)
    if [ "$Soft" ==  "yes" ]
    then
        printf "%s\n" "Start reset soft."
        git reset --soft HEAD~1
        printf "%s\n" "Start pushing data."
        git push -f
    fi
}

action_diff() {
    # Displays the uncommitted changes to the repository.
    git diff
}

action_reopen() {
    ./gitcommand.sh
}

action_help(){
    printf "\n%s\n" "status: helping"
    printf "\n%s" "Exit, Enter 0. "
    printf "\n%s" "Push, Enter 1. "
    printf "\n%s" "Pull, Enter 2. "
    printf "\n%s" "Remove Last Commit From Current Branch, Enter 3. "
    printf "\n%s" "Check Modified Differences, Enter 4. "
    printf "\n%s\n" "Reopen gitcommand.sh, Enter 9. "
    printf "\n"
}

hello_world() {

    while :
    do
        read -p "Enter the command: " -r Command

        case $Command in 
            0) see_you;;
            1) action_push;;
            2) action_pull;;
            3) action_soft;;
            4) action_diff;;
            help) action_help;;
            9) action_reopen;;

            *) printf "\n%s\n" "Command Error";;
        esac

        printf "\n%s" 
    done
}

see_you() {
    # Stop this program.
    printf "\n\n%s\n" "Till next time."
    printf "\n"
    exit 1
}

hello_world