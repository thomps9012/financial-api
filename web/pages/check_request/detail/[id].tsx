import Link from "next/link";
import styles from "../../../styles/Home.module.css";
import { GetServerSidePropsContext } from "next";

function CheckRequestDetail({ request_id }: { request_id: string }) {
  const approveRequest = async (e: any) => {
    const selected_permission = (
      document.getElementById("selected_permission") as HTMLSelectElement
    ).value;
    e.preventDefault();
  };
  const rejectRequest = async (e: any) => {
    e.preventDefault();
  };
  const archiveRequest = async (e: any) => {
    e.preventDefault();
  };
  return (
    <main >
      <h1>Check Request Detail for {request_id}</h1>
      <Link href={`/check_request/edit/${request_id}`}>Edit Request</Link>
      <button>Archive / Delete Request</button>
    </main>
  );
}
// possible phase back to getStaticProps
CheckRequestDetail.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};

export default CheckRequestDetail;
