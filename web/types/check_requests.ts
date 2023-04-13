import { Action } from "./actions";

export interface Purchase {
  amount: number;
  description: string;
  grant_line_item: string;
}

export interface Vendor {
  name: string;
  website?: string;
  address_line_one: string;
  address_line_two?: string;
}

export interface Check_Request {
  id: string;
  grant_id: string;
  user_id: string;
  date: string;
  category: string;
  vendor: Vendor;
  description: string;
  purchases: Purchase[];
  receipts: string[];
  order_total: number;
  credit_card: string;
  action_history: Action[];
  current_user: string;
  current_status: string;
  is_active: boolean;
}

export interface Check_Request_Overview {
  id: string;
  user_id: string;
  date: string;
  order_total: number;
  current_status: string;
  current_user: string;
  is_active: boolean;
}

export interface Check_Request_Input {
  grant_id: string;
  date: string;
  category: string;
  vendor: Vendor;
  description: string;
  purchases: Purchase[];
  receipts: string[];
  credit_card: string;
}

export interface Edit_Check_Request_Input {
  id: string;
  user_id: string;
  grant_id: string;
  date: string;
  category: string;
  vendor: Vendor;
  description: string;
  purchases: Purchase[];
  receipts: string[];
  credit_card: string;
}
