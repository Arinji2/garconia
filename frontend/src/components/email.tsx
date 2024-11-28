export default function Email() {
  return (
    <div className="shrink-0 w-[80%] h-[50px] flex flex-row items-center justify-center">
      <input
        className=" h-full bg-garconia-offwhite/50 shrink-0 w-[80%] outline-none rounded-l-md shadow-shadow border-[2px] border-black px-4"
        placeholder="Email..."
      />
      <button className="bg-garconia-red w-full  h-full rounded-r-md shadow-shadow border-[2px] border-black">
        <p className="text-white text-2xl font-bold"> NOTIFY ME!! </p>
      </button>
    </div>
  );
}
