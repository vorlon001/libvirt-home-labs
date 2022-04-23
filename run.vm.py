try:
    import CORE.LOADER
    CORE.LOADER.lm(globals(),
				"CORE.Nefelim", "sys"
				)

    CORE.LOADER.iglob(globals(),[
                            { "module": "CORE.Nefelim",   "method":"Nefelim",     "as": "Nefelim"    }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)


def main():

#    globals()["logger"] = initLog(loggingLevel = logging.DEBUG, exp = __name__)
#    print(sys.modules['__main__'].__dict__["logger"])

    nefelim = Nefelim()
    nefelim.Run()
    exit(1)

if __name__ == '__main__':
    main()

