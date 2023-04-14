import Navbar from "./navbar";
import Footer from "./footer";
import AccessDenied from "./accessDenied";
import { useAppContext } from "@/context/AppContext";
interface Props {
  children: React.ReactNode;
}

export default function Layout({ children }: Props) {
  const { logged_in } = useAppContext();
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
