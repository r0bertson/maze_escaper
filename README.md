# maze_escaper

### Finding a solution to a maze
Just pass your maze (as JSON) to `cmd/maze_escaper/main.go` like this:
```
cd cmd\maze_escaper
go run main.go "{\"forward\": \"exit\"}
```


### Generating a random maze and it's solution
You can simply run `cmd/generate_maze/main.go` with it's default params:
```
cd cmd\generate_maze
go run main.go
```

Or you can customize it passing the following arguments:

| Argument 	 | Default value           	 | Description                            	  |
|------------|---------------------------|-------------------------------------------|
| -d       	 | 10                      	 | Max depth of a path                    	  |
| -or      	 | 0.4                     	 | Obstacle occurrence rate                	 |
| -pr      	 | 0.3                     	 | Path's growth rate                     	  |
| -obs     	 | pgk/seed/obstacles.txt  	 | Path to obstacles seed file.           	  |
| -dir     	 | pgk/seed/directions.txt 	 | Path to directions file.               	  |
| -output  	 | false                   	 | Export maze and solution as json file. 	  |