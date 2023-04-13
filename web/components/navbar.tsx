// import { signOut, useSession } from "next-auth/react";
import Link from "next/link";
import { useState } from "react";
import styles from "./header.module.css";
import { useAppContext } from "@/context/AppContext";
import axios from "axios";

export default function Navbar() {
  // const { data: session } = useSession();
  const { login, logout } = useAppContext();
  const [openNav, setOpenNav] = useState(false);
  const { user_profile, logged_in } = useAppContext();
  return (
    <>
      <div className={styles.mobileHeader}>
        {logged_in ? (
          <>
            <Link href="/profile">
              <p>
                <span className={styles.signedInText}>
                  <br />
                  <strong>{user_profile.name}</strong>
                </span>
              </p>
            </Link>
            <p
              className={styles.menuOption}
              onClick={() => setOpenNav(!openNav)}
              style={{ fontSize: "35px" }}
            >
              Nav
            </p>
            <ul className={"navIcons-" + openNav}>
              <li className={styles.navIcon}>
                <Link href="/">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    🏡 Home
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/profile/inbox">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    📭 Inbox
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/mileage">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    🚗 Mileage
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/petty_cash">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    💵 Petty Cash
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/check_request">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    🗃️ Check Request
                  </p>
                </Link>
              </li>
              <li className={styles.navIcon}>
                <Link href="/how_to">
                  <p
                    className={styles.navIcon}
                    onClick={(e) => setOpenNav(false)}
                  >
                    🆘 Help / How To
                  </p>
                </Link>
              </li>
              {user_profile.admin && (
                <li className={styles.navIcon}>
                  <Link href="/users">
                    <p
                      className={styles.navIcon}
                      onClick={(e) => setOpenNav(false)}
                    >
                      👨‍👦‍👦 Users
                    </p>
                  </Link>
                </li>
              )}

              <li className={styles.navIcon}>
                <a
                  className={styles.navIcon}
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
          <ul className={"navIcons-" + openNav}>
            <li className={styles.navIcon}>
              <a
                className={styles.navIcon}
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
      </div>
      <header>
        <div className={styles.navHeader}>
          {logged_in ? (
            <>
              <Link href="/profile">
                <p>
                  <span className={styles.signedInText}>
                    <br />
                    <strong>{user_profile.name}</strong>
                  </span>
                </p>
              </Link>
              <ul className={styles.navIcons}>
                <li className={styles.navIcon}>
                  <Link href="/">
                    <p className={styles.navIcon}>
                      🏡<span className={styles.navSpan}>Home</span>
                    </p>
                  </Link>
                </li>
                <li className={styles.navIcon}>
                  <Link href="/profile/inbox">
                    <p className={styles.navIcon}>
                      📭<span className={styles.navSpan}>Inbox</span>
                    </p>
                  </Link>
                </li>
                <li className={styles.navIcon}>
                  <Link href="/mileage">
                    <p className={styles.navIcon}>
                      🚗<span className={styles.navSpan}>Mileage</span>
                    </p>
                  </Link>
                </li>
                <li className={styles.navIcon}>
                  <Link href="/petty_cash">
                    <p className={styles.navIcon}>
                      💵<span className={styles.navSpan}>Petty Cash</span>
                    </p>
                  </Link>
                </li>
                <li className={styles.navIcon}>
                  <Link href="/check_request">
                    <p className={styles.navIcon}>
                      🗃️<span className={styles.navSpan}>Check Request</span>
                    </p>
                  </Link>
                </li>
                <li className={styles.navIcon}>
                  <Link href="/how_to">
                    <p className={styles.navIcon}>
                      🆘<span className={styles.navSpan}>Help / How To</span>
                    </p>
                  </Link>
                </li>
                {user_profile.admin && (
                  <li className={styles.navIcon}>
                    <Link href="/users">
                      <p className={styles.navIcon}>
                        👨‍👦‍👦<span className={styles.navSpan}>Users</span>
                      </p>
                    </Link>
                  </li>
                )}
                <li className={styles.navIcon}>
                  <a
                    className={styles.navIcon}
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
            <ul className={styles.navIcons}>
              <li className={styles.navIcon}>
                <a
                  className={styles.navIcon}
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
        </div>
      </header>
    </>
  );
}
