import { title } from "process";
import * as React from "react";
import Project from "../Project";
import Section from "../Section";

const Achievements = (): React.ReactElement => {
  const sections = [
    {
      title: "College Related",
      projects: [
        <Project
          name="Ross 2"
          description="Contest Managing Platform"
          startYear="2021"
          endYear="2022"
          website="https://ross2.club"
          sourceCode="https://github.com/mbaraa/ross2"
          image="/ross2.png"
        />,
      ],
    },
  ];

  return (
    <div className="bg-[#2d333b] w-full h-full">
      {sections.map((s) => (
        <Section title={s.title} projects={s.projects} />
      ))}
    </div>
  );
};

export default Achievements;
