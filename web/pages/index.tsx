import { useAppContext } from "@/context/AppContext";
import styles from "@/styles/Home.module.css";
import Link from "next/link";

export default function Landing() {
  const { user_profile, incomplete_actions } = useAppContext();
  return (
    <main >
      <br />
      <header
        style={{
          display: "flex",
          justifyContent: "space-between",
          flexWrap: "wrap",
        }}
      >
        <h1>Financial Request Hub</h1>
        <Link href={"/profile/inbox"}>
          <p className="req-overview">
            {incomplete_actions.length} New Action Items
          </p>
        </Link>
      </header>
      <div className={styles.container}>
        <Link href={"/mileage"}>
          <h2>🚗 Mileage </h2>
        </Link>
        <hr />
        <Link href={"/mileage/create"}>
          <h3 style={{ fontWeight: 100 }}>New Request</h3>
        </Link>
        <br />
        <Link href={"/petty_cash"}>
          <h2>💵 Petty Cash </h2>
        </Link>
        <hr />
        <Link href={"/petty_cash/create"}>
          <h3 style={{ fontWeight: 100 }}>New Request</h3>
        </Link>
        <br />
        <Link href={"/check_request"}>
          <h2>🗃️ Check Requests </h2>
        </Link>
        <hr />
        <Link href={"/check_request/create"}>
          <h3 style={{ fontWeight: 100 }}>New Request</h3>
        </Link>
        <br />
        {user_profile.admin && (
          <>
            <h2>👨‍👦‍👦 Users </h2>
            <hr />
            <Link href={"/users"}>
              <h3 style={{ fontWeight: 100 }}>View All</h3>
            </Link>
          </>
        )}
      </div>
    </main>
  );
}
