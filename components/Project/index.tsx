import * as React from "react";
import Image from "next/image";
import Link from "next/link";

interface WSProps {
  name: string;
  link: string;
}

const Website = ({ name, link }: WSProps): React.ReactElement => {
  return (
    <div className="font-[SFUI] text-[16px] text-white inline-block underline">
      <Link href={link}>
        <a target="_blank">{name}</a>
      </Link>
    </div>
  );
};

interface Props {
  name: string;
  startYear: string;
  endYear?: string;
  description: string;
  website?: string;
  sourceCode?: string;
  image: string;
}

const Project = ({
  name,
  startYear,
  endYear,
  description,
  website,
  sourceCode,
  image,
}: Props): React.ReactElement => {
  return (
    <div className="text-white grid grid-cols-4 ml-[25px] sm:ml-[50px] mr-0 font-[SFUI] mt-[25px]">
      <div>
        <Image
          src={image}
          width={150}
          height={150}
          className="rounded-[100%]"
        />
      </div>
      <div className="pl-[15px] w-[250%]">
        <label className="italic text-[22px] font-bold">
          {name} ({startYear}-{endYear ?? "present"})
        </label>
        <div className="mt-[10px]" />
        <label className="w-[250%] text-[18px]">{description}</label>
        <div className="mt-[10px]" />
        {website && <Website name="Visit website..." link={website} />}
        &nbsp;
        {sourceCode && <Website name="View source code..." link={sourceCode} />}
      </div>
    </div>
  );
};

export default Project;
