import Navbar from "./navbar";
import Footer from "./footer";
// import { useSession } from "next-auth/react";
import AccessDenied from "./accessDenied";
// import Loading from "./loading";
import { useAppContext } from "@/context/AppContext";
interface Props {
  children: React.ReactNode;
}

export default function Layout({ children }: Props) {
  // const { data: session, status } = useSession();
  // const loading = status === "loading";
  const { logged_in } = useAppContext();
  // if (!session && loading) {
  // return <Loading />;
  // }
  // if (!session || !logged_in) {
  if (!logged_in) {
    return <AccessDenied />;
  }
  return (
    <>
      <Navbar />
      <main>{children}</main>
      <Footer />
    </>
  );
}
