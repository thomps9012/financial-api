import Error from "next/error";
import { useRouter } from "next/router";
// UNIMPLEMENTED
export default async function handleErrors({
  error,
  request,
  jwt,
}: {
  error: Error;
  request: Request;
  jwt: string;
}) {
  const formatted_error = {
    error_path: error,
    error_message: error,
  };
  const router = useRouter();
}
