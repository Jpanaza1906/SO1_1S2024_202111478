#include <linux/module.h>
#define INCLUDE_VERMAGIC
#include <linux/build-salt.h>
#include <linux/elfnote-lto.h>
#include <linux/export-internal.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

#ifdef CONFIG_UNWINDER_ORC
#include <asm/orc_header.h>
ORC_HEADER;
#endif

BUILD_SALT;
BUILD_LTO_INFO;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif



static const char ____versions[]
__used __section("__versions") =
	"\x14\x00\x00\x00\x7a\x67\xc1\x49"
	"single_open\0"
	"\x18\x00\x00\x00\xdb\x75\x43\x94"
	"_totalram_pages\0"
	"\x10\x00\x00\x00\xa6\x50\xba\x15"
	"jiffies\0"
	"\x1c\x00\x00\x00\x70\xfc\xbe\x37"
	"jiffies_to_msecs\0\0\0\0"
	"\x14\x00\x00\x00\x34\xe9\xb6\xe1"
	"init_task\0\0\0"
	"\x14\x00\x00\x00\xba\xcd\xf5\x73"
	"seq_printf\0\0"
	"\x1c\x00\x00\x00\xd8\x4e\x04\xb0"
	"remove_proc_entry\0\0\0"
	"\x14\x00\x00\x00\x21\xef\xe0\xda"
	"seq_read\0\0\0\0"
	"\x14\x00\x00\x00\xbb\x6d\xfb\xbd"
	"__fentry__\0\0"
	"\x14\x00\x00\x00\x26\x16\x07\x31"
	"proc_create\0"
	"\x10\x00\x00\x00\x7e\x3a\x2c\x12"
	"_printk\0"
	"\x1c\x00\x00\x00\xca\x39\x82\x5b"
	"__x86_return_thunk\0\0"
	"\x18\x00\x00\x00\x86\x05\x44\xe7"
	"module_layout\0\0\0"
	"\x00\x00\x00\x00\x00\x00\x00\x00";

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "CB12C4B4D64763574ECBBE7");
