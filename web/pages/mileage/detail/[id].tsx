import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function MileageRequestDetail({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Mileage Request Detail for {request_id}</h1>
      <Link href={`/mileage/edit/${request_id}`}>Edit Request</Link>
      <button>Archive / Delete Request</button>
    </main>
  );
}

MileageRequestDetail.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};

export default MileageRequestDetail;
