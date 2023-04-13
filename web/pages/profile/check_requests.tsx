import { useAppContext } from "@/context/AppContext";

export default function ProfileCheckPage() {
  const { user_profile } = useAppContext();
  const { id } = user_profile;
  return (
    <main>
      <h1>Check Request Page for {id}</h1>
    </main>
  );
}
