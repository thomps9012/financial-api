import styles from "@/styles/Home.module.css";
import Link from "next/link";

export default function Landing() {
  return (
    <main className={styles.landing}>
      <br />
      <header
        style={{
          display: "flex",
          justifyContent: "space-between",
          flexWrap: "wrap",
        }}
      >
        <h1>Financial Request Hub</h1>
        <Link href={"/me/inbox"}>
          {/* <a><p className='req-overview'>{notifications} New Action Items</p></a> */}
          <a>
            <p className="req-overview">0 New Action Items</p>
          </a>
        </Link>
      </header>
      <div className={styles.container}>
        <Link href={"/mileage"}>
          <a>
            <h2>🚗 Mileage </h2>
          </a>
        </Link>
        <hr />
        <Link href={"/mileage/create"}>
          <a>
            <h3 style={{ fontWeight: 100 }}>New Request</h3>
          </a>
        </Link>
        <br />
        <Link href={"/petty_cash"}>
          <a>
            <h2>💸 Petty Cash </h2>
          </a>
        </Link>
        <hr />
        <Link href={"/petty_cash/create"}>
          <a>
            <h3 style={{ fontWeight: 100 }}>New Request</h3>
          </a>
        </Link>
        <br />
        <Link href={"/check_request"}>
          <a>
            <h2>📑 Check Requests </h2>
          </a>
        </Link>
        <hr />
        <Link href={"/check_request/create"}>
          <a>
            <h3 style={{ fontWeight: 100 }}>New Request</h3>
          </a>
        </Link>
        <br />
        {/* {admin && <> */}
        <h2>👨‍👦‍👦 Users </h2>
        <hr />
        <Link href={"/users"}>
          <a>
            <h3 style={{ fontWeight: 100 }}>View All</h3>
          </a>
        </Link>
        {/* </> */}
        {/* } */}
      </div>
    </main>
  );
}
