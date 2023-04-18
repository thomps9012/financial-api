import { useAppContext } from "@/context/AppContext";
import styles from "@/styles/Home.module.css";
import Link from "next/link";

export default function Landing() {
  const { user_profile } = useAppContext();
  return (
    <main>
      <br />
      <header
        style={{
          display: "flex",
          justifyContent: "space-between",
          flexWrap: "wrap",
        }}
      >
        <h1>NORA Financial Requests</h1>
      </header>
      <section className={styles.home_page}>
        <Link href={"/mileage"}>
          <h2>Mileage 🚗</h2>
        </Link>
        <hr />
        <Link href={"/mileage/create"}>
          <h3>Create Request</h3>
        </Link>
        <Link href={"/profile/mileage"}>
          <h3>Your Mileage</h3>
        </Link>
        {user_profile.admin && (
          <>
            <Link href={"/mileage/reports/grant"}>
              <h3>Grant Report</h3>
            </Link>
            <Link href={"/mileage/reports/user"}>
              <h3>User Report</h3>
            </Link>
            <Link href={"/mileage/reports/monthly"}>
              <h3>Monthly Report</h3>
            </Link>
          </>
        )}
        <br />
        <Link href={"/petty_cash"}>
          <h2>Petty Cash 💵</h2>
        </Link>
        <hr />
        <Link href={"/petty_cash/create"}>
          <h3>Create Request</h3>
        </Link>
        <Link href={"/profile/petty_cash"}>
          <h3>Your Petty Cash</h3>
        </Link>
        {user_profile.admin && (
          <>
            <Link href={"/petty_cash/reports/grant"}>
              <h3>Grant Report</h3>
            </Link>
            <Link href={"/petty_cash/reports/user"}>
              <h3>User Report</h3>
            </Link>
            <Link href={"/petty_cash/reports/monthly"}>
              <h3>Monthly Report</h3>
            </Link>
          </>
        )}
        <br />
        <Link href={"/check_request"}>
          <h2>Check Requests 🗃️</h2>
        </Link>
        <hr />
        <Link href={"/check_request/create"}>
          <h3>Create Request</h3>
        </Link>
        <Link href={"/profile/check_requests"}>
          <h3>Your Check Requests</h3>
        </Link>
        {user_profile.admin && (
          <>
            <Link href={"/check_request/reports/grant"}>
              <h3>Grant Report</h3>
            </Link>
            <Link href={"/check_request/reports/user"}>
              <h3>User Report</h3>
            </Link>
            <Link href={"/check_request/reports/monthly"}>
              <h3>Monthly Report</h3>
            </Link>
          </>
        )}
        {user_profile.admin && (
          <>
            <br />
            <h2>Admin 🔐</h2>
            <hr />
            <Link href={"/users"}>
              <h3>View Users</h3>
            </Link>
          </>
        )}
      </section>
    </main>
  );
}
