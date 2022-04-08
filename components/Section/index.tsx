import * as React from "react";

interface Props {
  title: string;
  projects: React.ReactElement[];
}

const Section = ({ title, projects }: Props): React.ReactElement => {
  return (
    <div className="mt-[50px]">
      <h1 className="text-white font-[SFUI] ml-[15px] md:ml-[50px] text-[35px] italic text-bold">
        {title}
      </h1>
      {projects.map((p) => (
        <div key={Math.random()}>{p}</div>
      ))}
    </div>
  );
};

export default Section;
