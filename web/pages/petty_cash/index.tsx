import { useAppContext } from "@/context/AppContext";
import Link from "next/link";

export default function PettyCashOverview() {
  const { user_profile } = useAppContext();
  return (
    <main>
      <h1>Petty Cash</h1>
      <div className="hr" />
      <Link href="/petty_cash/create">
        <h2>Create New Request</h2>
      </Link>
      <Link href={"/profile/petty_cash"}>
        <h2>Your Active Requests</h2>
      </Link>
      {user_profile.admin && (
        <>
          <Link href="/petty_cash/reports/user">
            <h2>User Requests</h2>
          </Link>
          <Link href="/petty_cash/reports/grant">
            <h2>Requests by Grant</h2>
          </Link>
          <Link href="/petty_cash/reports/monthly">
            <h2>Monthly Requests</h2>
          </Link>
        </>
      )}
    </main>
  );
}
