try:
    import CORE.LOADER
    CORE.LOADER.lm(globals(), "logging","sys","uuid","inspect","traceback","typing", "os", "os.path")
    CORE.LOADER.iglob(globals(),[
                                  { "module": "typing",  "method":"Optional" }
                               ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)


def initLog(loggingLevel: int = logging.DEBUG, exp: str = __name__) -> Optional[logging.Logger]:
    try:
        logger = logging.getLogger(exp)
        logger.setLevel(loggingLevel)
        sh = logging.StreamHandler(sys.stdout)
        formatter = logging.Formatter('[%(asctime)s] %(levelname)s [%(filename)s.%(funcName)s:%(lineno)d] %(message)s', datefmt='%Y-%m-%d:%H:%M:%S')
        sh.setFormatter(formatter)
        logger.addHandler(sh)
        return logger
    except Exception as e:
        frame = inspect.currentframe()
        args, _, _, values = inspect.getargvalues(frame)
        print('function name "%s"' % inspect.getframeinfo(frame)[2])
        logging.critical({inspect.getframeinfo(frame)[2]: { i: values[i] for i in args}})
        logging.critical(e)
        return None

sys.modules['__main__'].__dict__["logger"] = initLog(loggingLevel = logging.DEBUG, exp = __name__)
sys.modules['__main__'].__dict__["DEBUG_MODE"] = False


def log(message: str = "", exception=None, exception_traceback=None):
#, type_message: str ="DEBUG", show_stacks: bool  = False, show_code_context: bool = False, exception=None, exception_traceback=None):

    logger = sys.modules['__main__'].__dict__["logger"]

    DEBUG_MODE = sys.modules['__main__'].DEBUG_MODE

    id_log = uuid.uuid4()
    stacks = inspect.stack()
    cf = inspect.currentframe()
    if DEBUG_MODE==True:
        logger.debug(f"LOGGER: uuid: {id_log}, {stacks[1][1]}:{cf.f_back.f_lineno}, message: {message}")
    else:
        logger.info(f"LOGGER: uuid: {id_log}, {stacks[1][1]}:{cf.f_back.f_lineno}, message: {message}")
    if exception is not None:
        #logger.debug(f"LOGGER: {{ uuid: {id_log}, globals: {exception.__traceback__.tb_frame.f_globals} }}")
        #logger.debug(f"LOGGER: {{ uuid: {id_log}, locals: {exception.__traceback__.tb_frame.f_locals} }}")
        logger.debug(f"LOGGER: {{ uuid: {id_log}, FileName: {exception.__traceback__.tb_frame.f_code.co_filename} }}")
        logger.debug(f"LOGGER: {{ uuid: {id_log}, lineno: {exception.__traceback__.tb_lineno} }}")
        f = exception.__traceback__.tb_frame
        h =  inspect.getframeinfo( f )
        args, _, _, values = inspect.getargvalues(f)
        stack_values = { "method": inspect.getframeinfo( f )[2], "args": { j: values[j] for j in args}, "variables": values }
        logger.debug(f"LOGGER: {{ uuid: {id_log}, stack values: {stack_values} }}")
    if DEBUG_MODE == True:
        for i in range(1,len(stacks)):
            stack = stacks[i]
            if DEBUG_MODE==True:
                logger.debug(f"STACK: {{ uuid: {id_log}, filename: {stack.filename}: function: {stack.function}: lineno:{stack.lineno}: code_context:{stack.code_context} }}")
            else:
                logger.debug(f"STACK: {{ uuid: {id_log}, filename: {stack.filename}: function: {stack.function}: lineno:{stack.lineno} }}")
        z = inspect.stack()
        for i in range(1,len(z)):
            f = z[i].frame
            h =  inspect.getframeinfo( f )
            args, _, _, values = inspect.getargvalues(f)
            stack_values = { "method": inspect.getframeinfo( f )[2], "args": { j: values[j] for j in args}, "variables": values }
            logger.debug(f"STACK: {{ uuid: {id_log}, stack values: {stack_values} }}")



def Dump( id, dump, from_class, f_code, stack):
    code = f_code.f_code
    print(f"DUMP {id}")
    print('\tModule/Class/Function : ' + code.co_filename + '::' + type(from_class).__name__ + '::' + code.co_name +'()')
    print('\tCalled from  Filename/Line/Method   : ' + os.path.basename(stack[1][1]) + '/' + str(stack[1][2]) + '/' + stack[1][3] + '()' )
    print(f"\tFILELINE        : {code.co_firstlineno}")
    print(f"\tARGUMENT COUNT  : {code.co_argcount}");
    print(f"\tLOCAL VARS      : {code.co_varnames}")
    print(f"\tDUMP            : {dump}")




