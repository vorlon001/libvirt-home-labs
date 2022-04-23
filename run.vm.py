try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
				"CORE.LibVirt.Nefelim", "sys"
				)

    boot.iglob(globals(),[
                            { "module": "CORE.LibVirt.Nefelim",   "method":"Nefelim",     "as": "Nefelim"    }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)


def main():

    nefelim = Nefelim()
    nefelim.Run()
    exit(1)

if __name__ == '__main__':
    main()

