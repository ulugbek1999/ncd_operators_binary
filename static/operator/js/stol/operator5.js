document.addEventListener("DOMContentLoaded", () => {

    // Function for handling format of the date
    function dateMyFormat(date) {
        var elem = date.split('.');
        var newMyDate = elem[2] + '-' + elem[1] + '-' + elem[0];
        //console.log(newMyDate);
        if(date){
            return newMyDate;
        }else{
            return '';
        }
    }

    var sendUniversal = document.querySelector("#send-universal")

    var clearData = [
        "#eduType",
        "#eduName",
        "#eduAddress",
        "#eduSpec",
        "#eduGradDate1",
        "#eduStartDate",
        "#eduAddit",
        "#eduFiles",
        "#langName",
        "#langLevel",
        "#langFiles",
        "#milName",
        "#milRegion",
        "#milSpecialization",
        "#milDateStarted",
        "#milDateGraduated",
        "#milRank",
        "#armyFiles",
        "#recName",
        "#recComment",
        "#rewFiles",
        "#marStatus",
        "#childrenAmount",
        "#familyFiles",
        "#expOrg",
        "#expDateStarted",
        "#expDateFinished",
        "#expPosition",
        "#expSupDibision",
        "#expAddress",
        "#expFiles",
    ]

    var clear = function () {
        clearData.forEach((element) => {
            if (document.querySelector(element)) {
                document.querySelector(element).value = ""
            }
        })
    }


    // Education 
    var eduCounter = 0
    var educations = []
    var eduPanel = document.querySelector("#edu-panel")
    const eduSave = document.querySelector("#eduSave")
    eduSave.addEventListener("click", () => {
        var eduType = document.querySelector("#eduType").value
        var eduName = document.querySelector("#eduName").value
        var eduAddress = document.querySelector("#eduAddress").value
        var eduSpec = document.querySelector("#eduSpec").value
        var eduGradDate = document.querySelector("#eduGradDate1").value
        var eduStartDate = document.querySelector("#eduStartDate").value
        var eduAddit = document.querySelector("#eduAddit").value
        var eduFiles = document.querySelector("#eduFiles")
        
        var eduData = {
            eduType: eduType,
            eduName: eduName,
            eduAddress: eduAddress,
            eduSpec: eduSpec,
            eduGradDate: dateMyFormat(eduGradDate),
            eduStartDate: dateMyFormat(eduStartDate),
            eduAddit: eduAddit,
            eduFiles: eduFiles.files[0]
        }
        console.log(dateMyFormat(eduGradDate))
        educations.push(eduData)
        const eduTemplate = `
<div class="op3View-item edu-all-${eduCounter}">
    <ul>
        <li>
            <strong>Тип/вид учебного заведения:</strong>
            <span class="edu-type-${eduCounter}>${educations[eduCounter].eduType}</span>
        </li>
        <li>
            <strong>Учебное заведение:</strong>
            <span class="edu-ins-${eduCounter}">${educations[eduCounter].eduName}</span>
        </li>
        <li>
            <strong>Адрес учебного заведения:</strong>
            <span class="edu-address-${eduCounter}">${educations[eduCounter].eduAddress}</span>
        </li>
        <li>
            <strong>Специальность:</strong>
            <span class="edu-specialty-${eduCounter}">${educations[eduCounter].eduSpec}</span>
        </li>
        <li>
            <strong>Дата начала обучения:</strong>
            <span class="edu-start-date-${eduCounter}">${educations[eduCounter].eduStartDate}</span>
        </li>
        <li>
            <strong>Дата окончания обучения:</strong>
            <span class="edu-end-date-${eduCounter}">${educations[eduCounter].eduGradDate}</span>
        </li>
        <li>
            <strong>Прочие данные об учебе:</strong>
            <span class="edu-other-${eduCounter}">${educations[eduCounter].eduAddit}</span>
        </li>
    </ul>
    <a href="" class="deleteButton eduDelete" data-index="${eduCounter}">Удалить</a>
</div>
    `
        var div = document.createElement("div")
        div.innerHTML = eduTemplate
        eduPanel.appendChild(div)
        eduCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-edu").style.display = "none"
        var eduDeletes = document.querySelectorAll('.eduDelete')
        eduDeletes.forEach((eduDelete) => {
            eduDelete.addEventListener('click', (event) => {
                event.preventDefault()
                document.querySelector(`.edu-all-${event.target.dataset.index}`).style.display = "none"
                educations.splice(event.target.dataset.index, 1)
            })
        })
    })

    // Language
    var langSave = document.querySelector("#lang-save")
    var languages = []
    var langCounter = 0
    var langPanel = document.querySelector("#lang-panel")
    langSave.addEventListener('click', () => {
        var langName = document.querySelector("#langName").value
        var langLevel = document.querySelector("#langLevel").value
        var langFiles = document.querySelector("#langFiles")

        var langData = {
            langName: langName,
            langLevel: langLevel,
            langFiles: langFiles.files[0]
        }
        languages.push(langData)
        const langTemplate = `
        <div class="op3View-item lang-all-${langCounter}">
            <ul>
                <li>
                    <strong>Название языка:</strong>
                    <span>${languages[langCounter].langName}</span>
                </li>
                <li>
                    <strong>Уровень знания:</strong>
                    <span>${languages[langCounter].langLevel}</span>
                </li>
            </ul>
            <a href="#" id="lang-save" data-index="${langCounter}" class="deleteButton langDelete">Удалить</a>
        </div>
    `
        var div = document.createElement("div")
        div.innerHTML = langTemplate
        langPanel.appendChild(div)
        langCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-lang").style.display = "none"
        var langDeletes = document.querySelectorAll(".langDelete")
        langDeletes.forEach((langDelete) => {
            langDelete.addEventListener("click", (event) => {
                event.preventDefault()
                document.querySelector(`.lang-all-${event.target.dataset.index}`).style.display = "none"
                languages.splice(event.target.dataset.index, 1)
            })
        })
    })

    // Army
    var armySave = document.querySelector("#army-save")
    var armies = []
    var armyCounter = 0
    var armyPanel = document.querySelector("#army-panel")
    armySave.addEventListener("click", () => {
        var milName = document.querySelector("#milName").value
        var milRegion = document.querySelector("#milRegion").value
        var milSpecialization = document.querySelector("#milSpecialization").value
        var milDateStarted = document.querySelector("#milDateStarted").value
        var milDateGraduated = document.querySelector("#milDateGraduated").value
        var milRank = document.querySelector("#milRank").value
        var armyFiles = document.querySelector("#armyFiles")

        var armyData = {
            milName: milName,
            milRegion: milRegion,
            milSpecialization: milSpecialization,
            milDateStarted: dateMyFormat(milDateStarted),
            milDateGraduated: dateMyFormat(milDateGraduated),
            milRank: milRank,
            armyFiles: armyFiles.files[0]
        }
        armies.push(armyData)
        const armyTemplate = `
        <div class="op3View-item army-all-${armyCounter}">
            <ul>
                <li>
                    <strong>Наименование воинской части:</strong>
                    <span>${armies[armyCounter].milName}</span>
                </li>
                <li>
                    <strong>Дата начала службы:</strong>
                    <span>${armies[armyCounter].milDateStarted}</span>
                </li>
                <li>
                    <strong>Регион:</strong>
                    <span>${armies[armyCounter].milRegion}</span>
                </li>
                <li>
                    <strong>Дата окончания службы:</strong>
                    <span>${armies[armyCounter].milDateGraduated}</span>
                </li>
                <li>
                    <strong>Специальность:</strong>
                    <span>${armies[armyCounter].milSpecialization}</span>
                </li>
                <li>
                    <strong>Звание:</strong>
                    <span>${armies[armyCounter].milRank}</span>
                </li>
            </ul>
            <a href="" class="deleteButton armyDelete" data-index=${armyCounter}>Удалить</a>
        </div>
        `
        var div = document.createElement("div")
        div.innerHTML = armyTemplate
        armyPanel.appendChild(div)
        armyCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-army").style.display = "none"
        document.querySelectorAll(".armyDelete").forEach((armyDelete) => {
            armyDelete.addEventListener("click", (event) => {
                event.preventDefault()
                document.querySelector(`.army-all-${event.target.dataset.index}`).style.display = "none"
                armies.splice(event.target.dataset.index, 1)
            })
        })
    })


    // Reward
    var rewSave = document.querySelector("#rew-save")
    var rewards = []
    var rewCounter = 0
    var rewPanel = document.querySelector("#rew-panel")
    rewSave.addEventListener('click', () => {
        var recName = document.querySelector("#recName").value
        var recComment = document.querySelector("#recComment").value
        var rewFiles = document.querySelector("#rewFiles")

        var rewData = {
            recName: recName,
            recComment: recComment,
            rewFiles: rewFiles.files[0]
        }
        rewards.push(rewData)
        const rewTemplate = `
        <div class="op3View-item rew-all-${rewCounter}">
            <ul>
                <li>
                    <strong>Название:</strong>
                    <span>${rewards[rewCounter].recName}</span>
                </li>
                <li>
                    <strong>Комментарий:</strong>
                    <span>${rewards[rewCounter].recComment}</span>
                </li>
            </ul>
            <a href="" class="deleteButton rewDelete" data-index="${rewCounter}">Удалить</a>
        </div>
        `
        var div = document.createElement("div")
        div.innerHTML = rewTemplate
        rewPanel.appendChild(div)
        rewCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-rew").style.display = "none"
        document.querySelectorAll(".rewDelete").forEach((rewDelete) => {
            rewDelete.addEventListener('click', (event) => {
                event.preventDefault()
                document.querySelector(`.rew-all-${event.target.dataset.index}`).style.display = "none"
                rewards.splice(event.target.dataset.index, 1)
            })
        })
    })

    // Family
    var familySave = document.querySelector("#family-save")
    var families = []
    var famCounter = 0
    var famPanel = document.querySelector("#family-panel")
    familySave.addEventListener("click", () => {
        var marStatus = document.querySelector("#marStatus").value
        var childrenAmount = document.querySelector("#childrenAmount").value
        var familyFiles = document.querySelector("#familyFiles")

        var famData = {
            marStatus: marStatus,
            childrenAmount: childrenAmount,
            familyFiles: familyFiles.files[0]
        }
        families.push(famData)
        const familyTemplate = `
        <div class="op3View-item family-all-${famCounter}">
            <ul>
                <li>
                    <strong>Семейный статус:</strong>
                    <span>${families[famCounter].marStatus}</span>
                </li>
                <li>
                    <strong>Количество детей:</strong>
                    <span>${families[famCounter].childrenAmount}</span>
                </li>
            </ul>
            <a href="" class="deleteButton familyDelete" data-index="${famCounter}">Удалить</a>
        </div>
        `
        var div = document.createElement("div")
        div.innerHTML = familyTemplate
        famPanel.appendChild(div)
        famCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-family").style.display = "none"
        document.querySelectorAll(".familyDelete").forEach((famDelete) => {
            famDelete.addEventListener('click', (event) => {
                event.preventDefault()
                document.querySelector(`.family-all-${event.target.dataset.index}`).style.display = "none"
                families.splice(event.target.dataset.index, 1)
            })
        })
    })

    // Experience
    var expSave = document.querySelector("#exp-save")
    var exps = []
    var expCounter = 0
    var expPanel = document.querySelector("#exp-panel")
    expSave.addEventListener("click", () => {
        var expOrg = document.querySelector("#expOrg").value
        var expDateStarted = document.querySelector("#expDateStarted").value
        var expDateFinished = document.querySelector("#expDateFinished").value
        var expPosition = document.querySelector("#expPosition").value
        var expSubDivision = document.querySelector("#expSupDibision").value
        var expAddress = document.querySelector("#expAddress").value
        var expFiles = document.querySelector("#expFiles")

        var expData = {
            expOrg: expOrg,
            expDateStarted: dateMyFormat(expDateStarted),
            expDateFinished: dateMyFormat(expDateFinished),
            expPosition: expPosition,
            expSubDivision: expSubDivision,
            expAddress: expAddress,
            expFiles: expFiles.files[0]
        }
        exps.push(expData)

        const expTempalate = `
        <div class="op3View-item exp-all-${expCounter}">
            <ul>
                <li>
                    <strong>Организация:</strong>
                    <span>${expOrg}</span>
                </li>
                <li>
                    <strong>Дата начала работы:</strong>
                    <span>${expDateStarted}</span>
                </li>
                <li>
                    <strong>Дата окончания работы:</strong>
                    <span>${expDateFinished}</span>
                </li>
                <li>
                    <strong>Должность:</strong>
                    <span>${expPosition}</span>
                </li>
                <li>
                    <strong>Подразделение:</strong>
                    <span>${expSubDivision}</span>
                </li>
                <li>
                    <strong>Место нахождения/адрес:</strong>
                    <span>${expAddress}</span>
                </li>
            </ul>
            <a href="" class="editButtonProfile expDelete" data-index="${expCounter}">Удалить</a>
        </div>
    `
        var div = document.createElement("div")
        div.innerHTML = expTempalate
        expPanel.appendChild(div)
        expCounter++
        clear()
        document.querySelector(".modal-activated").style.display = "none"
        document.querySelector("#operator-3-exp").style.display = "none"
        document.querySelectorAll(".expDelete").forEach((expDelete) => {
            expDelete.addEventListener('click', (event) => {
                event.preventDefault()
                document.querySelector(`.exp-all-${event.target.dataset.index}`).style.display = "none"
                exps.splice(event.target.dataset.index, 1)
            })
        })
    })

    sendUniversal.addEventListener("click", (event) => {
        const formData = new FormData()
        var fullNameEn = document.getElementById("fullNameEn").value
        var fullNameRu = document.getElementById("fullNameRu").value
        var passportSerial = document.getElementById("passportSerial").value
        var phone = document.getElementById("phone").value
        var birthDate = document.getElementById("birthDate").value
        var email = document.getElementById("email").value
        var passportExpires = document.getElementById("passportExpires").value
        var passportGivendate = document.getElementById("passportGivendate").value
        
        // Operator 1
        formData.append('full_name_en', fullNameEn)
        formData.append('full_name_ru', fullNameRu)
        formData.append('birth_date', dateMyFormat(birthDate))
        formData.append('phone', phone)
        formData.append('email', email)
        formData.append("gender", $("input[name='gender']:checked").val());
        formData.append('passport_serial', passportSerial)
        formData.append("passport_given_date", dateMyFormat(passportGivendate))
        formData.append("passport_expires", dateMyFormat(passportExpires))
        formData.append("inn", $("#inn").val());
        formData.append("birth_place_ru", $("#birthPlace").val());
        formData.append("living_address_ru", $("#livesAt").val());
        formData.append("register_number", $("#registerNumber").val());
        formData.append("group_id", $("#group-id").val());
        formData.append("send_sms", $("#sendSMS").is(":checked"));
        formData.append("send_email", $("#sendEmail").is(":checked"));
        formData.append("username", $("#login").val());
        formData.append("passport_image", $("input[name='passport_file']")[0].files[0]);

        // Operator 2
        formData.append("appearance", $("input[name='appearance']:checked").val());
        formData.append("neatness", $("input[name='neatness']:checked").val());
        formData.append("skin", $("input[name='skin']:checked").val());
        formData.append("height", $("#heightID").val());
        formData.append("weight", $("#weightID").val());
        formData.append("fatness", $("#fatnessID").val());
        formData.append("blood_group", $("input[name='blood_group']:checked").val());
        formData.append("blood_resus", $("input[name='blood_resus']:checked").val());
        formData.append("vision_l", $("#vision_lID").val());
        formData.append("vision_r", $("#vision_rID").val());
        formData.append("photo_1", $("input[name='photo_1']")[0].files[0]);
        formData.append("photo_2", $("input[name='photo_2']")[0].files[0]);
        formData.append("photo_3", $("input[name='photo_3']")[0].files[0]);
        formData.append("photo_4", $("input[name='photo_4']")[0].files[0]);

        // Operator 3
        // Education
        formData.append("edu_data", JSON.stringify(educations))
        if (educations.length >= 1) {
            educations.forEach((education) => {
                if (education.eduFiles != undefined) {
                    formData.append('edu_files', education.eduFiles)
                }
                else {
                    formData.append('edu_files', new File(["foo", "bar"], 'nil.txt'))
                }
            })
        }
        // Language
        formData.append("lang_data", JSON.stringify(languages))
        if (languages.length >= 1) {
            languages.forEach((language) => {
                if (language.langFiles != undefined) {
                    formData.append('lang_files', language.langFiles)
                }
                else {
                    formData.append('lang_files', new File(["foo", "bar"], "nil.txt"))
                }
            })
        }

        // Army
        formData.append("army_data", JSON.stringify(armies))
        if (armies.length >= 1) {
            armies.forEach((army) => {
                if (army.armyFiles != undefined) {
                    formData.append('army_files', army.armyFiles)
                }
                else {
                    formData.append('army_files', new File(["foo", "bar"], "nil.txt"))
                }
            })
        }

        // Reward
        formData.append("reward_data", JSON.stringify(rewards))
        if (rewards.length >= 1) {
            rewards.forEach((reward) => {
                if (reward.rewFiles != undefined) {
                    formData.append('reward_files', reward.rewFiles)
                }
                else {
                    formData.append('reward_files', new File(["foo", "bar"], "nil.txt"))
                }
            })
        }

        // Family
        formData.append("reward_data", JSON.stringify(families))
        if (families.length >= 1) {
            families.forEach((family) => {
                if (family.familyFiles != undefined) {
                    formData.append('family_files', reward.familyFiles)
                }
                else {
                    formData.append('family_files', new File(["foo", "bar"], "nil.txt"))
                }
            })
        }

        // Experience
        formData.append("exp_data", JSON.stringify(exps))
        if (exps.length >= 1) {
            exps.forEach((exp) => {
                if (exp.expFiles != undefined) {
                    formData.append('exp_files', exp.expFiles)
                }
                else {
                    formData.append('exp_files', new File(["foo", "bar"], "nil.txt"))
                }
            })
        }

        // Operator 5
        formData.append("wages", $("input[name='wages']:checked").val());
        formData.append("is_ready_for_university", $("input[name='to_university']:checked").val());
        formData.append("criminal_number", $("input[name='criminal_number']").val());
        formData.append("criminal_issue", $("input[name='criminal_issue']").val());
        formData.append("criminal_comment_ru", $("textarea[name='criminal_comment']").val());
        formData.append("add_spec_signs_ru", $("input[name='add_spec_signs']").val());
        formData.append("hobby_ru", $("input[name='hobby']").val());
        formData.append("other_ru", $("textarea[name='other']").val());
        formData.append("country", $("select[name='des_countries']").val());
        formData.append("psycholog", $("input[name='psixolog']:checked").val());
        formData.append("level", $("input[name='level']:checked").val());

        $.ajax({
            url: event.target.dataset.api,
            type: "POST",
            data: formData,
            contentType: false,
            processData: false,
            success() {
                alert("success")
            },
            error() {
                alert("error")
            }
        })
    
    })

})