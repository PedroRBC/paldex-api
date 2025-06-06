<h1 align="center">
  <img alt="Palworld Paldex" title="Palworld Paldex" src=".github/pal.png" width="200px" />
</h1>

<h3 align="center">
  Palworld Paldex API
</h3>
<h6 align="center"><i>Fast and lightweight Palworld API reimplementation in Go</i></h6>

<p align="center">
  <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/PedroRBC/paldex-api.svg">

  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/PedroRBC/paldex-api.svg">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/PedroRBC/paldex-api.svg">
  <a href="https://github.com/PedroRBC/paldex-api/commits/master">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/PedroRBC/paldex-api.svg">
  </a>

  <a href="https://github.com/PedroRBC/paldex-api/issues">
    <img alt="Repository issues" src="https://img.shields.io/github/issues/PedroRBC/paldex-api.svg">
  </a>

  <img alt="GitHub" src="https://img.shields.io/github/license/PedroRBC/paldex-api.svg">
</p>

<p align="center">
  <a href="#rocket-info">Info</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#computer-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#mag_right-functionalities">Functionalities</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#memo-license">License</a>
</p>

## :rocket: Info

This is the first Palworld API, it's a simple API to get all Palworld Paldex data. This is a fork of the original project for educational purposes.

<details>
  <summary>This is what final data looks like</summary>

```json
{
  "content": [
    {
      "id": 85,
      "key": "085",
      "image": "/public/images/paldeck/085.png",
      "name": "Relaxaurus",
      "wiki": "https://palworld.fandom.com/wiki/Relaxaurus",
      "types": ["dragon", "water"],
      "imageWiki": "https://static.wikia.nocookie.net/palworld/images/0/01/Relaxaurus_menu.png/",
      "suitability": [
        {
          "type": "watering",
          "image": "/public/images/works/watering.png",
          "level": 2
        },
        {
          "type": "transporting",
          "image": "/public/images/works/transporting.png",
          "level": 1
        }
      ],
      "drops": ["high_quality_pal_oil", "ruby"],
      "aura": {
        "name": "hungry_missile",
        "description": "Can be ridden. Can rapidly fire a missile launcher while mounted.",
        "tech": null
      },
      "description": "Contrary to its blasé appearance, it's quite ferocious.\nIt perceives everything in its sight as prey and will stop at nothing to devour it.",
      "skills": [
        {
          "level": 1,
          "name": "dragon_cannon",
          "type": "dragon",
          "cooldown": 2,
          "power": 30,
          "description": "Hurls an energy ball imbued with draconic energy at an enemy.\n"
        }
      ],
      "stats": {
        "hp": 110,
        "attack": {
          "melee": 110,
          "ranged": 100
        },
        "defense": 70,
        "speed": {
          "ride": 800,
          "run": 650,
          "walk": 60
        },
        "stamina": 100,
        "support": 100
      },
      "asset": "LazyDragon",
      "genus": "monster",
      "rarity": 8,
      "price": 10240,
      "size": "xl"
    }
  ],
  "page": 1,
  "limit": 10,
  "count": 1,
  "total": 1
}
```

</details>

## :computer: Technologies

This project was developed using the following technologies:

- [Go](https://golang.org/)
- [Vercel](https://vercel.com/)
- [VS Code](https://code.visualstudio.com/)
- [Docker](https://www.docker.com/)
- [Git](https://git-scm.com/)

## :mag_right: Functionalities

Current features:

- RESTful API endpoints for Palworld data
- Query parameter support for filtering results
- Docker support for easy deployment
- Go for high performance and reliability
- Clean code structure and organization

### Available Endpoints

- `GET /` - Get all pals
- `GET /:id` - Get pal by ID
- `GET /search/:name` - Search pals by name

### Query Parameters

| Param       | Type   | Description                                           |
| ----------- | ------ | ----------------------------------------------------- |
| id          | number | The ID of the Pal you want to get                     |
| name        | string | The name of the Pal you want to search                |

## :rocket: Related Projects

This project was inspired by and based on the structure and data from the original [Palworld Paldex API](https://github.com/mlg404/palworld-paldex-api) by Victor Eyer.

Other related projects that use the original API:

- [Paldex](https://github.com/viethua99/Paldex) – A Kotlin Multiplatform (KMM) app built with Compose Multiplatform to display Palworld game information.
- [PalBot](https://github.com/nibalizer/palbot-rs/) – A Discord bot that provides Palworld information using the Paldex API.
- [PalPad](https://github.com/Juanvic/PalPad) – A mobile app for browsing Palworld data.

---

## :star2: Credits

This project is a reimplementation of [palworld-paldex-api](https://github.com/mlg404/palworld-paldex-api) created by [Victor Eyer](https://github.com/mlg404).  
Special thanks to the original author for the amazing work and open data.

Original repository: [mlg404/palworld-paldex-api](https://github.com/mlg404/palworld-paldex-api)

---

## :memo: License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/PedroRBC/paldex-api/blob/master/LICENSE) file for more information.


---