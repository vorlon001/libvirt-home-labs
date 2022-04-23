try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(), "sys","inspect","typing","traceback","CORE.Core.Logger")
    boot.iglob(globals(),[
                                  { "module": "typing",  "method":"Dict"     },
                                  { "module": "typing",  "method":"TypeVar"  },
                                  { "module": "typing",  "method":"Callable" },
                                  { "module": "typing",  "method":"Optional" },
                                  { "module": "CORE.Core.Logger",      "method":"Dump"                   },
                                  { "module": "CORE.Core.Logger",      "method":"log"                    }
                               ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)



class Decorator(object):
    def __init__(self, arg):
        self.arg = arg
    def __call__(self, cls):
        class Wrapped(cls):
            def __setattr__(self, name, value):
                _ = Dump(id=41, dump={ "name": name, "value": value}, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                self.__dict__[name] = value

            def __getattr__(self, name):
                print(name)
                _ =Dump(id=51, dump={ "name": name }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                if name in self.__dict__['other_class']:
                    Dump(id=52, dump={ "name": name }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                    return self.__dict__['other_class'][name]
                elif name in self.__dict__:
                    Dump(id=52, dump={ "name": name }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                    return self.__dict__[name]
                return None

            def __getattribute__(self,name):
                attr = object.__getattribute__(self, name)
                if hasattr(attr, '__call__'):
                    def newfunc(*args, **kwargs):
                        Dump(id=31, dump={ "data": 'before calling %s' %attr.__name__ , "kwargs": kwargs }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None

                        result = attr(*args, **kwargs)

                        Dump(id=32, dump={ "result": result, "data": 'done calling %s' %attr.__name__ , "kwargs": kwargs}, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                        return result

                    return newfunc
                else:
                    return attr

            def __get__(self, instance, owner):
                Dump(id=21, dump={ "instance":instance, "owner": owner }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None
                return 'Экземпляр %s, класс %s' % (instance, owner)

            def __init__(self, **cls):
                self.other_class = cls
                super(Wrapped, self).__init__(**cls)
                Dump(id=100, dump={ "cls": self.other_class }, from_class=self, f_code=sys._getframe(), stack=inspect.stack()) if sys.modules['__main__'].DEBUG_MODE ==True else None

        return Wrapped

@Decorator("decorated class BASE")
class BASE(object):
    pass


@Decorator("decorated class WARP_DRIVE")
class WARP_DRIVE(BASE):

    @staticmethod
    def decorator(func):
        def tmp(*args, **kwargs):
          try:
            if len(kwargs)>0:
                for i in func.__annotations__:
                     if isinstance(kwargs[i],func.__annotations__[i])==False:
                         raise ValueError("Error type vars:{0} = {1} in args {0} = {2} '{3}'".format(i,func.__annotations__[i],type(kwargs[i]),kwargs[i]))
          except SystemExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 401, "kwargs":kwargs, "Exception": e }
          except KeyboardInterrupt as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 402, "kwargs":kwargs, "Exception": e }
          except GeneratorExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 403, "kwargs":kwargs, "Exception": e }
          except Exception as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 404, "kwargs":kwargs, "Exception": e }
          try:
            result = func(*args, **kwargs)
          except SystemExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 501, "kwargs":kwargs, "Exception": e }
          except KeyboardInterrupt as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 502, "kwargs":kwargs, "Exception": e }
          except GeneratorExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 503, "kwargs":kwargs, "Exception": e }
          except Exception as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 500, "kwargs":kwargs, "Exception": e }
          return { "code": 200, "kwargs":kwargs, "result": result }
        return tmp


    @staticmethod
    def decorator_void(func):
        def tmp(*args, **kwargs):
          try:
            if len(kwargs)>0:
                for i in func.__annotations__:
                     if isinstance(kwargs[i],func.__annotations__[i])==False:
                         raise ValueError("Error type vars:{0} = {1} in args {0} = {2} '{3}'".format(i,func.__annotations__[i],type(kwargs[i]),kwargs[i]))
          except SystemExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 401, "kwargs":kwargs, "Exception": e }
          except KeyboardInterrupt as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 402, "kwargs":kwargs, "Exception": e }
          except GeneratorExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 403, "kwargs":kwargs, "Exception": e }
          except Exception as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 404, "kwargs":kwargs, "Exception": e }
          try:
            func(*args, **kwargs)
          except SystemExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 501, "kwargs":kwargs, "Exception": e }
          except KeyboardInterrupt as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 502, "kwargs":kwargs, "Exception": e }
          except GeneratorExit as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 503, "kwargs":kwargs, "Exception": e }
          except Exception as e:
            log(message = "Error WARP_DRIVE {ERROR}".format(ERROR={"kwargs":kwargs, "Exception": None , "attr": func }))
            return { "code": 500, "kwargs":kwargs, "Exception": e }
          return { "code": 200, "kwargs":kwargs, "result": True }
        return tmp
