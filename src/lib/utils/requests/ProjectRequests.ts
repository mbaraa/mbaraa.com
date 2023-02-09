import type ProjectGroup from "$lib/models/ProjectGroup";
import Requests from "./Requests";

export default class ProjectRequests {
  static async getProjectGroups(): Promise<ProjectGroup[]> {
    return [
      {
        name: "Google Developer Student Clubs",
        description: "Stuff for GDSC - ASU to solve some problems",
        projects: [
          {
            name: "GDSC - ASU Home Page",
            description:
              "A home page for the GDSC - ASU chapter, that can be customized easily and reused for any other GDSC chapter.",
            startYear: "2022",
            endYear: "2023",
            website: "https://gdscasu.com",
            sourceCode: "https://github.com/GDSC-ASU/website",
            imagePath: `/images/gdg.png`,
          },
          {
            name: "GDSC Logo Generator",
            description:
              "My first web project, my Google Developer Student Clubs chapter's lead thought it would be a great idea if we had a logo generator that every other GDSC chapters can use it, so that every GDSC logos look the same in a neat way.",
            startYear: "2020",
            endYear: "2021",
            website: "https://logogen.gdscasu.com",
            sourceCode: "https://github.com/mbaraa/dsc_logo_generator",
            imagePath: `/images/logogen.png`,
          },
        ],
      },
      {
        name: "Gentoo Related",
        description: "Stuff that helps with Gentoo Linux",
        projects: [
          {
            name: "Eloi",
            description:
              "Gentoo ebuilds finder and installer, finds ebuilds from various overlay repos, and makes the needed cahnges to install an ebuild.",
            startYear: "2023",
            sourceCode: "https://github.com/mbaraa/eloi",
            imagePath: "/images/gentoo.svg",
          },
          {
            name: "Schwifter",
            description:
              "A Gentoo post installer script, inspired from Helmuthdu's AUI, [needs updates for new packages' versions]",
            startYear: "2019",
            endYear: "2019",
            sourceCode: "https://github.com/mbaraa/schwifter",
            imagePath: "/images/gentoo.svg",
          },
        ],
      },
      {
        name: "College Related",
        description: "Stuff for my faculty",
        projects: [
          {
            name: "Eurus",
            description:
              "The second enhanced version of Sheev, which will have full form automation, instead of just printing papers.",
            startYear: "2022",
            imagePath: `/images/eurus.png`,
            comingSoon: true,
          },
          {
            name: "Ross 2",
            description:
              "My biggest project yet, Ross is a university contest manager, it manages and automates all contest registration and closure routines.",
            startYear: "2021",
            endYear: "2022",
            website: "https://ross2.co",
            sourceCode: "https://github.com/mbaraa/ross2",
            imagePath: `/images/ross2.png`,
          },
          {
            name: "Sheev",
            description:
              "Form to image genrator, I made this project because of the lack of digitized forms in my university.",
            startYear: "2021",
            endYear: "2021",
            website: "https://sheev.vercel.app",
            sourceCode: "https://github.com/mbaraa/sheev",
            imagePath: `/images/sheev.png`,
          },
        ],
      },
      {
        name: "Misc Web",
        description: "Some other web stuff",
        projects: [
          {
            name: "Ladder and Snake",
            description: "A weird looking Ladder and Snake game.",
            startYear: "2022",
            endYear: "2022",
            sourceCode: "https://github.com/mbaraa/ladder_and_snake",
            imagePath: "/images/ladder-and-snake.png",
          },
          {
            name: "Temco MEP Home Page",
            description:
              "A home page for Temco MEP that displays the services and featurs the company offers.",
            startYear: "2022",
            endYear: "2022",
            sourceCode: "",
            website: "https://temco-mep.com",
            imagePath: "/images/temco-mep.png",
          },
          {
            name: "Shorts Ninja",
            description:
              "My second web project, I was exploring web and I decided to go with the classic hello web project i.e. a URL Shortner.",
            startYear: "2021",
            endYear: "2021",
            sourceCode: "https://github.com/mbaraa/shortsninja",
            imagePath: `/images/shortsninja.png`,
          },
        ],
      },
      {
        name: "Terminal Games",
        description: "Some games to play in the terminal",
        projects: [
          {
            name: "Snek",
            description:
              "Funny story, I saw a snake screen saver, and thought to myself, it would be great if I made a snake game, soon it'll solve itself!",
            startYear: "2022",
            endYear: "2022",
            sourceCode:
              "https://github.com/mbaraa/console_games/tree/master/Snek",
            imagePath: `/images/snek.png`,
          },
          {
            name: "Tic Tac Toe",
            description: "I was boared again :)",
            startYear: "2021",
            endYear: "2021",
            sourceCode:
              "https://github.com/mbaraa/console_games/tree/master/TicTacToe",
            imagePath: `/images/ttt.png`,
          },
          {
            name: "Tetris",
            description:
              "Terminal based tetris game, this is my fist Go project ever, I made it because I had nothing else to do.",
            startYear: "2020",
            endYear: "2020",
            sourceCode:
              "https://github.com/mbaraa/console_games/tree/master/TheTetrisProject",
            imagePath: `/images/tetris.png`,
          },
        ],
      },
    ];
  }
}
