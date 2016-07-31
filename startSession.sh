#!/bin/bash
tmux new-session -d -s go
tmux rename-window goEnv
tmux split-window -h
tmux resize-pane -t 0 -R 15 
tmux attach-session -t go
