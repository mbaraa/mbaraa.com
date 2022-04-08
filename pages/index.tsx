import type { NextPage } from "next";
import Header from "../components/Header";
import Name from "../components/Name";
import Achievements from "../components/Achievements";
import Separator from "../components/Separator";

const Home: NextPage = () => {
  return (
    <div className="font-[SFUI] text-[20px]">
      <Header />
      <Separator />
      <Name />
      <Separator />
      <Achievements />
    </div>
  );
};

export default Home;
