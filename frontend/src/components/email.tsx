export default function Email() {
  return (
    <div className="shrink-0 w-[80%] h-fit xl:h-[50px] gap-2 xl:gap-0 flex flex-col xl:flex-row items-center justify-center">
      <input
        type="email"
        className=" xl:h-full text-garconia-white h-[50px] bg-garconia-offwhite/50 xl:shrink-0 w-full xl:rounded-none xl:w-[80%] outline-none rounded-md xl:rounded-l-md shadow-shadow border-[2px] border-black px-4"
        placeholder="Email..."
      />
      <button className="px-2 bg-garconia-red w-fit xl:w-full h-[40px] xl:h-full rounded-md xl:rounded-none xl:rounded-r-md shadow-shadow border-[2px] border-black">
        <p className="text-white whitespace-nowrap  md:text-2xl font-bold">
          {" "}
          NOTIFY ME!!{" "}
        </p>
      </button>
    </div>
  );
}
