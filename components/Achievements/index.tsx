import * as React from "react";
import Project from "../Project";
import Section from "../Section";

const Achievements = (): React.ReactElement => {
  const sections = [
    {
      title: "College Related",
      projects: [
        <Project
          key={Math.random()}
          name="Ross 2"
          description="My biggest project ever, Ross is a university contest manager, it manages and automates all contest registration and closure routines."
          startYear="2021"
          endYear="2022"
          website="https://ross2.club"
          sourceCode="https://github.com/mbaraa/ross2"
          image="/ross2.png"
        />,
        <Project
          key={Math.random()}
          name="Sheev"
          description="Form to image genrator, I made this project because of the lack of digitalized forms in my university, so it can help everyone fill form easily, without having to waith for other people to approve the forms or not."
          startYear="2021"
          endYear="2022"
          website="https://sheev.vercel.app"
          sourceCode="https://github.com/mbaraa/sheev"
          image="/sheev.png"
        />,
      ],
    },
    {
      title: "Early Web",
      projects: [
        <Project
          key={Math.random()}
          name="Shorts Ninja"
          description="My second web project, I was exploring web and I decided to go with the classic hello web project i.e. a URL Shortner"
          startYear="2020"
          endYear="2021"
          website="https://shorts.ninja"
          sourceCode="https://github.com/mbaraa/shortsninja"
          image="/shortsninja.png"
        />,
        <Project
          key={Math.random()}
          name="GDSC Logo Generator"
          description="My first web project, my Google Developer Student Clubs chapter's lead thought it would be a great idea if we had a logo generator that every other GDSC chapters can use it, so that every GDSC logos look the same in a neat way."
          startYear="2020"
          endYear="2021"
          website="https://logogen.dscasu.com"
          sourceCode="https://github.com/mbaraa/dsc_logo_generator"
          image="/gdg.png"
        />,
      ],
    },
    {
      title: "Terminal Games",
      projects: [
        <Project
          key={Math.random()}
          name="Snek"
          description="Funny story, I saw a snake screen saver, and thought to myself, it would be great if I made a snake game, soon it'll solve itself!"
          startYear="2022"
          endYear="2022"
          sourceCode="https://github.com/mbaraa/console_games/tree/master/Snek"
          image="/snek.png"
        />,
        <Project
          key={Math.random()}
          name="Tic Tac Toe"
          description="I was boared again :)"
          startYear="2021"
          endYear="2021"
          sourceCode="https://github.com/mbaraa/console_games/tree/master/TicTacToe"
          image="/ttt.png"
        />,
        <Project
          key={Math.random()}
          name="Tetris"
          description="Terminal based tetris game, this is my fist Go project ever, I made it because I had nothing else to do."
          startYear="2020"
          endYear="2020"
          sourceCode="https://github.com/mbaraa/console_games/tree/master/TheTetrisProject"
          image="/tetris.png"
        />,
      ],
    },
  ];

  return (
    <div className="bg-[#2d333b] w-full h-full">
      <h1 className="text-white font-[SFUI] ml-[20px] md:ml-[70px] mb-[-20px] mt-[20px] text-[40px] italic text-bold">
        My Stuff...
      </h1>
      {sections.map((s) => (
        <div key={Math.random()}>
          <Section title={s.title} projects={s.projects} />
        </div>
      ))}
    </div>
  );
};

export default Achievements;
