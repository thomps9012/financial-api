import Link from "next/link";
import { useState } from "react";
import styles from "./header.module.css";
import { useAppContext } from "@/context/AppContext";

export default function Navbar() {
  const { login, logout } = useAppContext();
  const [openNav, setOpenNav] = useState(false);
  const { user_profile, logged_in } = useAppContext();
  return (
    <>
      <nav className={styles.mobileHeader}>
        {logged_in ? (
          <>
            <h1
              className={styles.menuOption}
              onClick={() => setOpenNav(true)}
              style={{ fontSize: "25px" }}
            >
              Open Nav
            </h1>
            <ul className={"navIcons-" + openNav}>
              <li className={styles.navIcon}>
                <a
                  style={{ position: "absolute", right: 10 }}
                  onClick={() => setOpenNav(false)}
                >
                  Hide Menu ❌
                </a>
              </li>
              <li className={styles.navIcon}>
                <Link href="/profile">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    {user_profile.name} Profile
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    🏡 Homepage
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/profile/inbox">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    📭 Incomplete Actions
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/mileage">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    🚗 Mileage Requests
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/petty_cash">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    💵 Petty Cash Requests
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/check_request">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    🗃️ Check Requests
                  </p>
                </Link>
              </li>
              {user_profile.admin && (
                <li className={styles.navIcon}>
                  <Link href="/users">
                    <p
                      className={styles.navIcon}
                      onClick={() => setOpenNav(false)}
                    >
                      👨‍👦‍👦 Users
                    </p>
                  </Link>
                </li>
              )}
              <li className={styles.navIcon}>
                <Link href="/how_to">
                  <p
                    className={styles.navIcon}
                    onClick={() => setOpenNav(false)}
                  >
                    🆘 Help & How To
                  </p>
                </Link>
              </li>
            </ul>
            <ul className={styles.login}>
              <li className={styles.navIcon}>
                <a
                  onClick={(e: any) => {
                    e.preventDefault();
                    logout();
                  }}
                >
                  Sign Out 🚀
                </a>
              </li>
            </ul>
          </>
        ) : (
          <ul className={styles.login}>
            <li className={styles.navIcon}>
              <a
                onClick={(e: any) => {
                  e.preventDefault();
                  const data = {
                    id: "109157735191825776845",
                    name: "TEST FINANCE",
                    email: "test@example.com",
                  };
                  login(data);
                }}
              >
                🚀 Sign In
              </a>
            </li>
          </ul>
        )}
      </nav>
      <nav className={styles.navHeader}>
        {logged_in ? (
          <>
            <ul className={styles.navIcons}>
              <li className={styles.navIcon}>
                <Link href="/">
                  <p className={styles.navIcon}>Home</p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/profile">
                  <p className={styles.navIcon}>{user_profile.name} Profile</p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/mileage">
                  <p className={styles.navIcon}>Mileage Requests</p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/petty_cash">
                  <p className={styles.navIcon}>Petty Cash Requests</p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/check_request">
                  <p className={styles.navIcon}>Check Requests</p>
                </Link>
              </li>
              {user_profile.admin && (
                <li className={styles.navIcon}>
                  <Link href="/users">
                    <p className={styles.navIcon}>Users</p>
                  </Link>
                </li>
              )}
              <li className={styles.navIcon}>
                <Link href="/how_to">
                  <p className={styles.navIcon}>Help & How To</p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <a
                  className={styles.navIcon}
                  style={{ right: 20, top: 40, position: "absolute" }}
                  onClick={(e: any) => {
                    e.preventDefault();
                    logout();
                  }}
                >
                  Sign Out 🚀
                </a>
              </li>
            </ul>
          </>
        ) : (
          <ul className={styles.login}>
            <li className={styles.navIcon}>
              <a
                onClick={(e: any) => {
                  e.preventDefault();
                  const data = {
                    id: "109157735191825776845",
                    name: "TEST FINANCE",
                    email: "test@example.com",
                  };
                  login(data);
                }}
              >
                🚀 Sign In
              </a>
            </li>
          </ul>
        )}
      </nav>
    </>
  );
}
