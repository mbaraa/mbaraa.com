import * as React from "react";
import Link from "next/link";

interface SLProps {
  name: string;
  link: string;
}

const SocialLink = ({ name, link }: SLProps): React.ReactElement => {
  return (
    <div className="italic font-[SFUI] text-[20px] text-white font-bold inline-block underline">
      <Link href={link}>
        <a target="_blank">{name}</a>
      </Link>
    </div>
  );
};

const Header = (): React.ReactElement => {
  return (
    <div className="bg-[#2d333b] h-[100px]">
      <div className="absolute right-1 mt-[35px] mr-[20px]">
        <SocialLink name="GitHub" link="https://github.com/mbaraa" />
        <label className="text-white">
          {"  "}•{"  "}
        </label>
        <SocialLink name="Twitter" link="https://twitter.com/mbaraa271" />
      </div>
    </div>
  );
};

export default Header;
