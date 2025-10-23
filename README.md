# pokedex-cli

A command-line Pokédex written in Go.  
Explore the Pokémon world, discover locations, catch Pokémon, view your Pokédex entries, and display caught Pokémon as terminal ASCII art.

This project is originally based on a Boot.dev assignment, expanded with additional functionality such as:
- ASCII sprite rendering
- Future plans for battle simulation

---

## Features

- View map locations (`map`, `mapb`)
- Explore areas to find available Pokémon (`explore <location>`)
- Catch Pokémon by name (`catch <pokemon>`)
- See your captured Pokémon list (`pokedex`)
- Inspect a captured Pokémon's data (`inspect <pokemon>`)
- Display a captured Pokémon as ASCII art (`image <pokemon>`)
- Built-in help and graceful exit

Example:
```bash
> map
> explore viridian-forest
> catch pikachu
> image pikachu
```

---

## Commands

| Command              | Parameters    | Description                      |
| -------------------- | ------------- | -------------------------------- |
| `help`               | —             | Display help                     |
| `exit`               | —             | Quit the Pokédex                 |
| `map`                | —             | Show next 20 regions             |
| `mapb`               | —             | Show previous 20 regions         |
| `explore <location>` | location name | Show Pokémon found in that area  |
| `catch <pokemon>`    | Pokémon name  | Attempt to catch a Pokémon       |
| `inspect <pokemon>`  | Pokémon name  | View details of a caught Pokémon |
| `pokedex`            | —             | List all caught Pokémon          |
| `image <pokemon>`    | Pokémon name  | Render captured Pokémon as ASCII |

---

## Installation

```bash
go install github.com/ethan-mdev/pokemon-cli@latest
```

Then run:

```bash
pokedex
```

---

## Dependencies

ASCII image rendering powered by:

* [https://github.com/TheZoraiz/ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter)
  Licensed under Apache License 2.0

Pokémon data and sprite URLs from:

* [https://pokeapi.co](https://pokeapi.co)

---

## Planned Features

* Pokémon battle system
* Smarter catching logic (stats + difficulty)
* Saving/loading Pokédex progress across sessions

---

## Legal Notice

This is an **unofficial fan project**.

Pokémon and Pokémon character names, images, and assets are trademarks of
Nintendo, Game Freak, Creatures Inc., and The Pokémon Company.

* No copyrighted Pokémon images are stored in this repository.
* Pokémon sprites and data are retrieved at runtime from PokéAPI.
* This project is non-commercial and educational only.

---

## License

All original code in this repository is licensed under the **MIT License**.
See [`LICENSE`](./LICENSE) for details.

