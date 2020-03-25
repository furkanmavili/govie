# Govie

Govie is a CLI app for storing movies/series lists.

## Installation



```bash
git clone https://github.com/furkanmavili/govie && cd govie/bin
```

## Usage
Show all commands with help:
```bash
./govie help
```

Show all lists:
```bash
./govie list all
```
Create list:
```bash
./govie create list_name
```
Add movies/series to list:
```bash
./govie add --list=list_name --rate=10 the godfather
```

Show the contents of the list:
```bash
./govie show list_name
```
Delete movies/series from list:
```bash
./govie delete --list=list_name the godfather
```




## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
