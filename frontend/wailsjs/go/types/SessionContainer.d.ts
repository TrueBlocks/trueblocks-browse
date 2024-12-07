// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';
import {context} from '../models';

export function Accumulate(arg1:types.Nothing):Promise<void>;

export function CleanWindowSize(arg1:context.Context):Promise<types.Window>;

export function Clear():Promise<void>;

export function CollateAndFilter(arg1:types.Filter):Promise<any>;

export function Finalize():Promise<void>;

export function ForEveryItem(arg1:types.EveryNothingFn,arg2:any):Promise<boolean>;

export function GetAddress():Promise<string>;

export function GetChain():Promise<string>;

export function GetFile():Promise<string>;

export function GetFolder():Promise<string>;

export function GetItems():Promise<any>;

export function GetRoute():Promise<string>;

export function GetTab(arg1:string):Promise<string>;

export function GetWindow():Promise<types.Window>;

export function GetWizardStr():Promise<string>;

export function IsFlagOn(arg1:string):Promise<boolean>;

export function Load():Promise<void>;

export function MarshalJSON():Promise<Array<number>>;

export function NeedsUpdate():Promise<boolean>;

export function Save():Promise<void>;

export function SetAddress(arg1:string):Promise<void>;

export function SetChain(arg1:string):Promise<void>;

export function SetFile(arg1:string):Promise<void>;

export function SetFlagOn(arg1:string,arg2:boolean):Promise<void>;

export function SetFolder(arg1:string):Promise<void>;

export function SetItems(arg1:any):Promise<void>;

export function SetRoute(arg1:string):Promise<void>;

export function SetTab(arg1:string,arg2:string):Promise<void>;

export function SetWindow(arg1:types.Window):Promise<void>;

export function SetWizardStr(arg1:string):Promise<void>;

export function ShallowCopy():Promise<types.Containerer>;

export function Sort():Promise<void>;

export function String():Promise<string>;

export function UnmarshalJSON(arg1:Array<number>):Promise<void>;
