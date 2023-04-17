import Navbar from "./navbar";
import Footer from "./footer";
import AccessDenied from "./accessDenied";
import { useAppContext } from "@/context/AppContext";
import { useRouter } from "next/router";
import Head from "next/head";
import { useEffect, useState } from "react";
import routeTitle from "@/utils/routeHandler";
interface Props {
  children: React.ReactNode;
}

export default function Layout({ children }: Props) {
  const { logged_in } = useAppContext();
  const router = useRouter();
  const [title, setTitle] = useState("NORA | Finance Requests");
  useEffect(() => {
    const page_title = routeTitle(router.route);
    setTitle(page_title);
  }, [router.route]);
  // if (!logged_in) {
  //   return (
  //     <>
  //       <Head>
  //         <title>{title}</title>
  //         <meta property="og:type" content="website" />
  //         <meta name="og:title" property="og:title" content={title} />
  //         <link rel="preconnect" href="https://fonts.googleapis.com" />
  //         <link
  //           rel="preconnect"
  //           href="https://fonts.gstatic.com"
  //           crossOrigin=""
  //         />
  //         <link
  //           href="https://fonts.googleapis.com/css2?family=Cabin&family=Catamaran&family=Jost&family=Overpass&family=Quicksand&family=Raleway&family=Tajawal&family=Urbanist&display=swap"
  //           rel="stylesheet"
  //         ></link>
  //       </Head>
  //       <AccessDenied />;
  //       <Footer />
  //     </>
  //   );
  // }
  return (
    <>
      <Head>
        <title>{title}</title>
        <meta property="og:type" content="website" />
        <meta name="og:title" property="og:title" content={title} />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin=""
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Cabin&family=Catamaran&family=Jost&family=Overpass&family=Quicksand&family=Raleway&family=Tajawal&family=Urbanist&display=swap"
          rel="stylesheet"
        ></link>
      </Head>
      <Navbar />
      <main>{children}</main>
      <Footer />
    </>
  );
}
