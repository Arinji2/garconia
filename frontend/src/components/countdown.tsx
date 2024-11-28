import Countdown from "react-countdown";

export default function CountDownComponent({ date }: { date: Date }) {
  return (
    <div className="w-fit h-fit">
      <Countdown date={date} renderer={CountDownRender} />,
    </div>
  );

  function CountDownRender({
    seconds,
    minutes,
    hours,
    days,
  }: {
    seconds: number;
    minutes: number;
    hours: number;
    days: number;
  }) {
    return (
      <section className="grid grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-9 items-center justify-center w-fit h-fit">
        <Count time={days} label="DAYS" />
        <Count time={hours} label="HOURS" />
        <Count time={minutes} label="MINUTES" />
        {/* Center the SECONDS component when there are 3 columns */}
        <div className="md:col-span-3 xl:col-span-1 flex justify-center">
          <Count time={seconds} label="SECONDS" />
        </div>
      </section>
    );
  }

  function Count({ time, label }: { time: number; label: string }) {
    const digits = String(time).padStart(2, "0").split("").map(Number);
    return (
      <div className="flex flex-row gap-6 items-center justify-center w-fit h-fit">
        <div className="w-fit h-fit flex flex-col items-center justify-center gap-3">
          <div className="w-fit h-fit flex flex-row items-center justify-center gap-3">
            <div className="bg-garconia-red w-[60px] flex flex-col items-center justify-center h-fit rounded-lg shadow-shadow border-[2px] border-black p-4">
              <p className="text-garconia-white text-xl md:text-2xl font-bold">
                {digits[0]}
              </p>
            </div>
            <div className="bg-garconia-red w-[60px] flex flex-col items-center justify-center h-fit rounded-lg shadow-shadow border-[2px] border-black p-4">
              <p className="text-garconia-white text-xl md:text-2xl font-bold">
                {digits[1]}
              </p>
            </div>
          </div>
          <p className="text-garconia-white text-lg md:text-2xl font-bold uppercase">
            {label}
          </p>
        </div>
      </div>
    );
  }
}
