import { useAppContext } from "@/context/AppContext";

export default function ProfileCheckPage() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Check Request Page for {name}</h1>
    </main>
  );
}
