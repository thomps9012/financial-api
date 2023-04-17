import { useAppContext } from "@/context/AppContext";
import Link from "next/link";

export default function MileageOverview() {
  const { user_profile } = useAppContext();
  return (
    <main>
      <h1>Mileage Requests</h1>
      <div className="hr" />
      <Link id="new" href="/mileage/create">
        <p>Create New</p>
      </Link>
      <Link href={"/profile/mileage"}>
        <h2>Your Active Requests</h2>
      </Link>
      {user_profile.admin && (
        <>
          <Link href="/mileage/reports/user">
            <h2>User Requests</h2>
          </Link>
          <Link href="/mileage/reports/grant">
            <h2>Requests by Grant</h2>
          </Link>
          <Link href="/mileage/reports/monthly">
            <h2>Monthly Requests</h2>
          </Link>
        </>
      )}
    </main>
  );
}
