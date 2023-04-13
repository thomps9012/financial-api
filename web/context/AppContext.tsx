import { Incomplete_Action } from "@/types/actions";
import { Grant } from "@/types/grants";
import {
  Axios_Credentials,
  Login_Input,
  User_Name_Info,
  User_Public_Info,
} from "@/types/users";
import {
  Dispatch,
  ReactNode,
  SetStateAction,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import axios from "axios";
import { Vehicle } from "@/types/users";
import { Mileage_Overview } from "@/types/mileage";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { Check_Request_Overview } from "@/types/check_requests";
import { useRouter } from "next/router";
type Props = {
  children: ReactNode;
};

type App_Context_Type = {
  user_credentials: Axios_Credentials;
  logged_in: boolean;
  user_profile: User_Public_Info;
  user_list: User_Name_Info[];
  grant_list: Grant[];
  incomplete_actions: Incomplete_Action[];
  setActions: Dispatch<SetStateAction<Incomplete_Action[]>>;
  clearAction: (action_id: string) => void;
  login: (login_info: Login_Input) => void;
  logout: () => void;
};

const default_app_context: App_Context_Type = {
  user_credentials: {
    headers: {
      Authorization: "",
    },
    withCredentials: true,
  },
  logged_in: false,
  user_profile: {
    id: "",
    email: "",
    name: "",
    last_login: new Date(),
    vehicles: new Array<Vehicle>(),
    is_active: false,
    admin: false,
    permissions: [""],
    incomplete_actions: new Array<Incomplete_Action>(),
    mileage_requests: new Array<Mileage_Overview>(),
    petty_cash_requests: new Array<Petty_Cash_Overview>(),
    check_requests: new Array<Check_Request_Overview>(),
  },
  user_list: [],
  grant_list: [],
  incomplete_actions: [],
  setActions: () => {},
  clearAction: () => {},
  login: () => {},
  logout: () => {},
};

const AppContext = createContext<App_Context_Type>(default_app_context);

export function useAppContext() {
  return useContext(AppContext);
}

export function AppProvider({ children }: Props) {
  const router = useRouter();
  useEffect(() => {
    setTimeout(logout_user, 1000 * 60 * 60 * 12);
  }, [router.route]);
  const grants = [
    {
      id: "H79TI082369",
      name: "BCORR",
    },
    {
      id: "SOR_HOUSING",
      name: "SOR Recovery Housing",
    },
    {
      id: "SOR_PEER",
      name: "SOR Peer",
    },
    {
      id: "SOR_LORAIN",
      name: "SOR Lorain 2.0",
    },
    {
      id: "H79TI085495",
      name: "RAP AID (Recover from Addition to Prevent Aids)",
    },
    {
      id: "2020-JY-FX-0014",
      name: "JSBT (OJJDP) - Jumpstart For A Better Tomorrow",
    },
    {
      id: "H79SP082264",
      name: "HIV Navigator",
    },
    {
      id: "H79SP082475",
      name: "SPF (HOPE 1)",
    },
    {
      id: "SOR_TWR",
      name: "SOR 2.0 - Together We Rise",
    },
    {
      id: "H79TI083370",
      name: "BSW (Bridge to Success Workforce)",
    },
    {
      id: "H79SM085150",
      name: "CCBHC",
    },
    {
      id: "H79TI083662",
      name: "IOP New Syrenity Intensive outpatient Program",
    },
    {
      id: "TANF",
      name: "TANF",
    },
    {
      id: "H79SP081048",
      name: "STOP Grant",
    },
    {
      id: "H79TI085410",
      name: "N MAT (NORA Medication-Assisted Treatment Program)",
    },
  ];
  const user_list = [{ id: "", name: "", active: false }];
  const [user_token, setToken] = useState("");
  const [user_logged_in, setLoggedIn] = useState(false);
  const [user_info, setUserInfo] = useState({
    id: "",
    email: "",
    name: "",
    last_login: new Date(),
    vehicles: new Array<Vehicle>(),
    is_active: false,
    admin: false,
    permissions: [""],
    incomplete_actions: new Array<Incomplete_Action>(),
    mileage_requests: new Array<Mileage_Overview>(),
    petty_cash_requests: new Array<Petty_Cash_Overview>(),
    check_requests: new Array<Check_Request_Overview>(),
  });
  const [incompleteActions, setIncompleteActions] = useState(
    new Array<Incomplete_Action>()
  );
  const clearAction = (action_id: string) => {
    const new_state = incompleteActions.filter(({ id }) => id != action_id);
    setIncompleteActions(new_state);
  };
  const setProfileInfo = async (auth_token: string) => {
    try {
      const { data } = await axios.get("/api/me", {
        headers: {
          Authorization: `Bearer ${auth_token}`,
        },
        withCredentials: true,
      });
      console.log(data.data);
      setUserInfo(data.data);
    } catch (error) {
      console.error(error);
    }
  };
  const login_user = async (user_info: Login_Input) => {
    try {
      const res = await axios.post("/api/auth/login", user_info);
      console.log(res);
      if ((res && res.data?.data && res.status === 200) || res.status === 201) {
        setLoggedIn(true);
        setToken(res.data.data.token);
        await setProfileInfo(res.data.data.token);
      }
    } catch (error) {
      console.error(error);
    }
  };
  const logout_user = async () => {
    try {
      const res = await axios.post("/api/auth/logout");
      console.log(res);
      if (res.status === 200) {
        setLoggedIn(false);
        setToken("");
        setUserInfo({
          id: "",
          email: "",
          name: "",
          last_login: new Date(),
          vehicles: new Array<Vehicle>(),
          is_active: false,
          admin: false,
          permissions: [""],
          incomplete_actions: new Array<Incomplete_Action>(),
          mileage_requests: new Array<Mileage_Overview>(),
          petty_cash_requests: new Array<Petty_Cash_Overview>(),
          check_requests: new Array<Check_Request_Overview>(),
        });
        window.location.assign("/");
      }
    } catch (error) {
      console.error(error);
    }
  };
  const value = {
    user_profile: user_info,
    user_credentials: {
      headers: {
        Authorization: `Bearer ${user_token}`,
      },
      withCredentials: true,
    },
    logged_in: user_logged_in,
    user_list: user_list,
    grant_list: grants,
    incomplete_actions: incompleteActions,
    setActions: setIncompleteActions,
    clearAction: clearAction,
    login: login_user,
    logout: logout_user,
  };

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
}
