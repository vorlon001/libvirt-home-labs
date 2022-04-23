try:
    def lm(glob, *name):
       for i in name:
           if i in globals():
               glob[i] = globals()[i]
           else:
               importlib = __import__("importlib.util")
               importlib_util=__import__("importlib.util")
               found = importlib.util.find_spec(i) is not None
               if found==True:
                   try:
                       globals()[i] = importlib.import_module(i)
                       glob[i] = globals()[i]
                   except Exception as e:
                       import traceback
                       print("INTERNAL ERROR CORE:",traceback.format_exc());
                       raise ImportError('Error Load Module STEP 1 {}'.format(i))
               else:
                   import traceback
                   print("INTERNAL ERROR CORE:",traceback.format_exc());
                   raise ImportError('Error Load Module not found STEP 2 "{}"'.format(i))
       return True

    def warp_drive(func):
        def tmp(*args, **kwargs):
            try:
                for i in func.__annotations__:
                    if isinstance(kwargs[i],func.__annotations__[i])==False:
                        raise ValueError("Error type vars:{0} = {1} in args {0} = {2} '{3}'".format(i,func.__annotations__[i],type(kwargs[i]),kwargs[i]))
            except BaseException as e:
                import traceback
                print("INTERNAL ERROR CORE:",traceback.format_exc());
                return { "code": 404, "kwargs":kwargs, "Exception": e }, True
            try:
                result = func(*args, **kwargs)
            except BaseException as e:
                import traceback
                print("INTERNAL ERROR CORE:",traceback.format_exc());
                return { "code": 500, "kwargs":kwargs, "Exception": e },True
            return { "code": 200, "kwargs":kwargs, "result": result }, False
        return tmp

    iglob = lambda glob, m: { "a": glob.update({ v['as'] if 'as' in v else v['method']: glob[v['module']].__dict__[v['method']]}) for v in m }
except Exception as e:
    print("type error: " + str(e),__file__)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)

