// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';
import {base} from '../models';
import {menu} from '../models';
import {context} from '../models';
import {config} from '../models';
import {wizard} from '../models';
import {app} from '../models';
import {output} from '../models';

export function AbiPage(arg1:number,arg2:number):Promise<types.AbiContainer>;

export function AddrToName(arg1:base.Address):Promise<string>;

export function CancelAllContexts():Promise<void>;

export function CancelContext(arg1:base.Address):Promise<void>;

export function ConvertToAddress(arg1:string):Promise<base.Address|boolean>;

export function ExportToCsv(arg1:base.Address):Promise<void>;

export function FileNew(arg1:menu.CallbackData):Promise<void>;

export function FileOpen(arg1:menu.CallbackData):Promise<void>;

export function FileSave(arg1:menu.CallbackData):Promise<void>;

export function FileSaveAs(arg1:menu.CallbackData):Promise<void>;

export function GetChain():Promise<string>;

export function GetChainInfo(arg1:string):Promise<types.Chain>;

export function GetChainList():Promise<Array<string>>;

export function GetContext():Promise<context.Context>;

export function GetDaemonJson(arg1:string):Promise<string>;

export function GetEnv(arg1:string):Promise<string>;

export function GetExploreUrl(arg1:string,arg2:boolean,arg3:boolean):Promise<string>;

export function GetLastAddress():Promise<base.Address>;

export function GetMenus():Promise<menu.Menu>;

export function GetMeta():Promise<types.MetaData>;

export function GetSession():Promise<config.Session>;

export function GetSessionDeamon(arg1:string):Promise<boolean>;

export function GetSessionSubVal(arg1:string):Promise<string>;

export function GetSessionVal(arg1:string):Promise<string>;

export function GetSessionWizard():Promise<wizard.State>;

export function GetWindow():Promise<config.Window>;

export function GetWizardState():Promise<wizard.State>;

export function HelpToggle(arg1:menu.CallbackData):Promise<void>;

export function HistoryPage(arg1:string,arg2:number,arg3:number):Promise<types.HistoryContainer>;

export function IndexPage(arg1:number,arg2:number):Promise<types.IndexContainer>;

export function IsConfigured():Promise<boolean>;

export function ManifestPage(arg1:number,arg2:number):Promise<types.ManifestContainer>;

export function ModifyAbi(arg1:app.ModifyData):Promise<void>;

export function ModifyMonitors(arg1:app.ModifyData):Promise<void>;

export function ModifyName(arg1:app.ModifyData):Promise<void>;

export function ModifyNoop(arg1:app.ModifyData):Promise<void>;

export function ModifyProject(arg1:app.ModifyData):Promise<void>;

export function MonitorPage(arg1:number,arg2:number):Promise<types.MonitorContainer>;

export function NamePage(arg1:number,arg2:number):Promise<types.NameContainer>;

export function Navigate(arg1:string,arg2:string):Promise<void>;

export function ProjectPage(arg1:number,arg2:number):Promise<types.ProjectContainer>;

export function Refresh(arg1:Array<string>):Promise<void>;

export function RegisterCtx(arg1:base.Address):Promise<output.RenderCtx>;

export function Reload(arg1:base.Address):Promise<void>;

export function SetChain(arg1:string,arg2:base.Address):Promise<void>;

export function SetEnv(arg1:string,arg2:string):Promise<void>;

export function SetSessionDaemon(arg1:string,arg2:boolean):Promise<void>;

export function SetSessionVal(arg1:string,arg2:string):Promise<void>;

export function StateToString(arg1:string):Promise<string>;

export function StatusPage(arg1:number,arg2:number):Promise<types.StatusContainer>;

export function StepWizard(arg1:wizard.Step):Promise<wizard.State>;

export function String():Promise<string>;

export function SystemAbout(arg1:menu.CallbackData):Promise<void>;

export function SystemQuit(arg1:menu.CallbackData):Promise<void>;

export function ToggleDaemon(arg1:string):Promise<void>;

export function ViewAbis(arg1:menu.CallbackData):Promise<void>;

export function ViewDaemons(arg1:menu.CallbackData):Promise<void>;

export function ViewHistory(arg1:menu.CallbackData):Promise<void>;

export function ViewIndexes(arg1:menu.CallbackData):Promise<void>;

export function ViewManifest(arg1:menu.CallbackData):Promise<void>;

export function ViewMonitors(arg1:menu.CallbackData):Promise<void>;

export function ViewNames(arg1:menu.CallbackData):Promise<void>;

export function ViewProject(arg1:menu.CallbackData):Promise<void>;

export function ViewSettings(arg1:menu.CallbackData):Promise<void>;

export function ViewStatus(arg1:menu.CallbackData):Promise<void>;

export function ViewWizard(arg1:menu.CallbackData):Promise<void>;
