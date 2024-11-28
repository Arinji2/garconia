export default function Socials() {
  return (
    <div className="w-fit h-fit flex flex-col items-stretch gap-3 justify-center">
      <p className="text-white text-xl font-bold">Follow Our Socials</p>
      <div className="w-fit flex flex-row items-stretch h-auto justify-center">
        <div className=" h-[100px] w-fit flex flex-col items-end justify-center">
          <img className="w-fit h-fit" src="/flag.png" />
        </div>
        <div className="bg-garconia-offwhite mx-3  h-full w-[2px]"></div>
        <div className=" gap-4 h-full w-fit flex flex-col items-start justify-center">
          <a
            href=""
            className="w-fit flex flex-row items-center justify-center h-fit gap-2"
          >
            <img width={30} height={30} src="/discord-color.svg" />
            <p className="text-garconia-white text-xl font-medium">Discord</p>
          </a>
          <a
            href=""
            className="w-fit flex flex-row items-center justify-center h-fit gap-2"
          >
            <img width={30} height={30} src="/bluesky-color.svg" />
            <p className="text-garconia-white text-xl font-medium">Bluesky</p>
          </a>
        </div>
      </div>
    </div>
  );
}
