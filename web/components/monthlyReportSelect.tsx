export default function MonthlyReportSelect({
  reportType,
  setReport,
}: {
  reportType: string;
  setReport: any;
}) {
  const handleSubmit = (e: any) => {
    e.preventDefault();
    const date_input = (
      document.getElementById("month_select") as HTMLInputElement
    ).value.split("-");
    const month = parseInt(date_input[1]);
    const year = parseInt(date_input[0]);
    let handleSubmit = (e: any) => {
      e.preventDefault();
      switch (reportType.trim().toLowerCase().split(" ").join("_")) {
        case "check":
          break;
        case "mileage":
          break;
        case "petty_cash":
          break;
        default:
          break;
      }
    };
  };
  return (
    <form>
      <label htmlFor="month_select">
        {reportType} Report (month and year):
      </label>
      <input type="month" id="month_select" name="month_select" />
      <a onClick={handleSubmit} className="archive-btn">
        {" "}
        Generate Report{" "}
      </a>
    </form>
  );
}
