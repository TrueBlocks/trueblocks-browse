// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {base} from '../models';
import {menu} from '../models';
import {types} from '../models';
import {context} from '../models';
import {daemons} from '../models';
import {config} from '../models';
import {output} from '../models';

export function AddrToName(arg1:base.Address):Promise<string>;

export function Cancel(arg1:base.Address):Promise<number|boolean>;

export function ConvertToAddress(arg1:string):Promise<base.Address|boolean>;

export function Fatal(arg1:string):Promise<void>;

export function FileNew(arg1:menu.CallbackData):Promise<void>;

export function FileOpen(arg1:menu.CallbackData):Promise<void>;

export function FileSave(arg1:menu.CallbackData):Promise<void>;

export function FileSaveAs(arg1:menu.CallbackData):Promise<void>;

export function Freshen(arg1:Array<string>):Promise<void>;

export function GetAbis(arg1:number,arg2:number):Promise<types.AbiContainer>;

export function GetAbisCnt():Promise<number>;

export function GetContext():Promise<context.Context>;

export function GetDaemon(arg1:string):Promise<daemons.Daemon>;

export function GetExistingAddrs():Promise<Array<string>>;

export function GetHistory(arg1:string,arg2:number,arg3:number):Promise<types.TransactionContainer>;

export function GetHistoryCnt(arg1:string):Promise<number>;

export function GetIndex(arg1:number,arg2:number):Promise<types.IndexContainer>;

export function GetIndexCnt():Promise<number>;

export function GetLast(arg1:string):Promise<string>;

export function GetLastDaemon(arg1:string):Promise<boolean>;

export function GetManifest(arg1:number,arg2:number):Promise<types.ManifestContainer>;

export function GetManifestCnt():Promise<number>;

export function GetMenus():Promise<menu.Menu>;

export function GetMonitors(arg1:number,arg2:number):Promise<types.MonitorContainer>;

export function GetMonitorsCnt():Promise<number>;

export function GetNames(arg1:number,arg2:number):Promise<types.NameContainer>;

export function GetNamesCnt():Promise<number>;

export function GetSession():Promise<config.Session>;

export function GetStatus(arg1:number,arg2:number):Promise<types.StatusContainer>;

export function GetStatusCnt():Promise<number>;

export function HelpToggle(arg1:menu.CallbackData):Promise<void>;

export function RegisterCtx(arg1:base.Address):Promise<output.RenderCtx>;

export function SetLast(arg1:string,arg2:string):Promise<void>;

export function SetLastDaemon(arg1:string,arg2:boolean):Promise<void>;

export function StateToString(arg1:string):Promise<string>;

export function String():Promise<string>;

export function SystemAbout(arg1:menu.CallbackData):Promise<void>;

export function SystemQuit(arg1:menu.CallbackData):Promise<void>;

export function ToggleDaemon(arg1:string):Promise<void>;

export function ViewAbis(arg1:menu.CallbackData):Promise<void>;

export function ViewDaemons(arg1:menu.CallbackData):Promise<void>;

export function ViewHistory(arg1:menu.CallbackData):Promise<void>;

export function ViewHome(arg1:menu.CallbackData):Promise<void>;

export function ViewIndexes(arg1:menu.CallbackData):Promise<void>;

export function ViewManifest(arg1:menu.CallbackData):Promise<void>;

export function ViewMonitors(arg1:menu.CallbackData):Promise<void>;

export function ViewNames(arg1:menu.CallbackData):Promise<void>;

export function ViewSettings(arg1:menu.CallbackData):Promise<void>;

export function ViewStatus(arg1:menu.CallbackData):Promise<void>;
