import { useAppContext } from "@/context/AppContext";
import Link from "next/link";

export default function CheckRequestOverview() {
  const { user_profile } = useAppContext();
  return (
    <main>
      <h1>Check Requests</h1>
      <div className="hr" />
      <Link id="new" href="/check_request/create">
        <p>Create New</p>
      </Link>
      <Link href={"/profile/check_request"}>
        <h2>Your Active Requests</h2>
      </Link>
      {user_profile.admin && (
        <>
          <Link href="/check_request/reports/user">
            <h2>User Requests</h2>
          </Link>
          <Link href="/check_request/reports/grant">
            <h2>Requests by Grant</h2>
          </Link>
          <Link href="/check_request/reports/monthly">
            <h2>Monthly Requests</h2>
          </Link>
        </>
      )}
    </main>
  );
}
