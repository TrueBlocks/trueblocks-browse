// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {base} from '../models';
import {menu} from '../models';
import {types} from '../models';
import {servers} from '../models';
import {config} from '../models';
import {output} from '../models';

export function Cancel(arg1:base.Address):Promise<number|boolean>;

export function ConvertToAddress(arg1:string):Promise<base.Address|boolean>;

export function Fatal(arg1:string):Promise<void>;

export function FileNew(arg1:menu.CallbackData):Promise<void>;

export function FileOpen(arg1:menu.CallbackData):Promise<void>;

export function FileSave(arg1:menu.CallbackData):Promise<void>;

export function FileSaveAs(arg1:menu.CallbackData):Promise<void>;

export function GetAbisCnt():Promise<number>;

export function GetAbis(arg1:number,arg2:number):Promise<types.AbiSummary>;

export function GetExistingAddrs():Promise<Array<string>>;

export function GetHistoryCnt(arg1:string):Promise<number>;

export function GetHistory(arg1:string,arg2:number,arg3:number):Promise<Array<types.TransactionEx>>;

export function GetIndex(arg1:number,arg2:number):Promise<types.IndexSummary>;

export function GetIndexCnt():Promise<number>;

export function GetLast(arg1:string):Promise<string>;

export function GetManifest(arg1:number,arg2:number):Promise<types.ManifestSummary>;

export function GetManifestCnt():Promise<number>;

export function GetMenus():Promise<menu.Menu>;

export function GetMonitorsCnt():Promise<number>;

export function GetMonitors(arg1:number,arg2:number):Promise<types.MonitorSummary>;

export function GetNamesCnt():Promise<number>;

export function GetNames(arg1:number,arg2:number):Promise<types.NameSummary>;

export function GetServer(arg1:string):Promise<servers.Server>;

export function GetSession():Promise<config.Session>;

export function GetStatus(arg1:number,arg2:number):Promise<types.StatusEx>;

export function GetStatusCnt():Promise<number>;

export function HelpToggle(arg1:menu.CallbackData):Promise<void>;

export function RegisterCtx(arg1:base.Address):Promise<output.RenderCtx>;

export function SetLast(arg1:string,arg2:string):Promise<void>;

export function StartServers():Promise<void>;

export function StateToString(arg1:string):Promise<string>;

export function String():Promise<string>;

export function SystemAbout(arg1:menu.CallbackData):Promise<void>;

export function SystemQuit(arg1:menu.CallbackData):Promise<void>;

export function ToggleServer(arg1:string):Promise<void>;

export function ViewAbis(arg1:menu.CallbackData):Promise<void>;

export function ViewHistory(arg1:menu.CallbackData):Promise<void>;

export function ViewHome(arg1:menu.CallbackData):Promise<void>;

export function ViewIndexes(arg1:menu.CallbackData):Promise<void>;

export function ViewManifest(arg1:menu.CallbackData):Promise<void>;

export function ViewMonitors(arg1:menu.CallbackData):Promise<void>;

export function ViewNames(arg1:menu.CallbackData):Promise<void>;

export function ViewServers(arg1:menu.CallbackData):Promise<void>;

export function ViewSettings(arg1:menu.CallbackData):Promise<void>;

export function ViewStatus(arg1:menu.CallbackData):Promise<void>;
